package client

import (
	"bytes"
	"fmt"
	"github.com/practice/multi_cluster_client/pkg/model"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net"
	"time"
)

// sSHConnect 使用 ssh 登入
func sSHConnect(user, password, host string, port string) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallback,
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%v", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

// GetClientByRemoteKubeConfig 从远程服务器获取 kubeconfig path，
// 并解析，获取到客户端。
// 输入：kubeconfig：远端 kubeconfig 目录 insecure：是否跳过证书
func GetClientByRemoteKubeConfig(remoteNode *model.RemoteNode, kubeconfigPath string, insecure bool) (*kubernetes.Clientset, error) {
	session, err := sSHConnect(remoteNode.User, remoteNode.Password, remoteNode.Host, remoteNode.Port)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
	session.Stderr = &stdErr

	output, err := session.Output(fmt.Sprintf("cat %v", kubeconfigPath))
	if err != nil {
		return nil, err
	}

	if stdErr.String() != "" {
		fmt.Println("exec result get error")
		fmt.Println(string(stdErr.Bytes()))
	} else {
		fmt.Println(string(stdOut.Bytes()))
	}
	// 加载 kubeconfig 对象
	config, err := clientcmd.Load(output)
	// config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	// 将 kubeconfig 中的集群配置分配给 *rest.Config
	restConfig, err := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	restConfig.Insecure = insecure
	// 使用 *rest.Config 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
