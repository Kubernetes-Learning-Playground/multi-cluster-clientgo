package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	//"github.com/go-yaml/yaml"
	"io/ioutil"
	//"k8s.io/apimachinery/pkg/util/json"
	//"k8s.io/client-go/tools/clientcmd/api"
	"log"
)


func main() {
	// 读取 kubeconfig 文件内容
	kubeconfigPath := "/Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config1"
	data, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(data))
	config, err := clientcmd.Load(data)
	fmt.Println("err3: ", err)
	//config, err := clientcmd.LoadFromFile(kubeconfigPath)
	// 将 kubeconfig 中的集群配置分配给 rest.Config
	restConfig, err := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	restConfig.Insecure = true
	fmt.Println("err2: ", err)
	if err != nil {
		panic(err.Error())
	}
	// 使用 rest.Config 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(restConfig)
	fmt.Println("err1: ", err)
	if err != nil {
		panic(err.Error())
	}
	p1 , err := clientset.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	fmt.Println("err: ", err)
	for _, v := range p1.Items {
		fmt.Println(v.Name)
		fmt.Println("aaa")
	}
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(config)
	// 解析 kubeconfig 文件
	//var config KubeConfig
	//ccc := &KubeConfig{}
	//var AAA map[string]interface{}
	//
	//err = json.Unmarshal(data, &AAA)
	//if err != nil {
	//	log.Fatal("aaa", err)
	//}
	//
	//fmt.Println(AAA)

	// 可以使用解析后的 KubeConfig 对象进行操作
	//log.Printf("API Version: %s\n", ccc.APIVersion)
	//log.Printf("Kind: %s\n", ccc.Kind)
	//log.Printf("Current Context: %s\n", ccc.CurrentContext)
	// ...
}