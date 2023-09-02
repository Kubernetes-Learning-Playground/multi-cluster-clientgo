package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"context"
	"log"
	"os/exec"
)


func main() {
	// 远程服务器的 SSH 信息
	sshUser := "root"
	sshHost := "42.193.17.123"
	sshPort := "22"
	sshKeyPath := "~/.ssh/k8s_use1.pem"

	// 要读取的文件路径
	filePath := ".kube/config"

	// 构建 SSH 命令
	sshCmd := exec.Command("ssh",
		"-i", sshKeyPath,
		"-p", sshPort,
		"-o", "StrictHostKeyChecking=no",
		fmt.Sprintf("%s@%s", sshUser, sshHost),
		"cat", filePath,
	)

	//// 将命令输出连接到标准输出
	//sshCmd.Stdout = os.Stdout
	//
	//// 执行 SSH 命令
	//err := sshCmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 执行 SSH 命令并获取输出
	output, err := sshCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

	config, err := clientcmd.Load(output)
	// config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	// 将 kubeconfig 中的集群配置分配给 rest.Config
	restConfig, err := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	fmt.Println("err2: ", err)
	if err != nil {
		log.Fatal(err)
	}
	//restConfig.Insecure = true
	// 使用 rest.Config 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(restConfig)
	fmt.Println("err1: ", err)
	if err != nil {
		log.Fatal(err)
	}
	p1 , err := clientset.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	fmt.Println("err: ", err)
	for _, v := range p1.Items {
		fmt.Println(v.Name)
		fmt.Println("aaa")
	}

}
