package main

import (
	"context"
	"fmt"
	"golanglearning/new_project/multi_cluster_client/pkg/client"
	"golanglearning/new_project/multi_cluster_client/pkg/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// 1. 创建文件
	c, _ := config.BuildConfig("./config.yaml")
	multiClient, _ := client.NewForConfig(c)
	// 2. 使用时与原生client-go几乎相同，只是需要多指定集群名，
	p1, _ := multiClient.Cluster("cluster1").CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	for _, v := range p1.Items {
		fmt.Println(v.Name)
	}

	p2, _ := multiClient.Cluster("cluster2").CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	for _, v := range p2.Items {
		fmt.Println(v.Name)
	}

	p3, _ := multiClient.Cluster("cluster3").CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	for _, v := range p3.Items {
		fmt.Println(v.Name)
	}

	//p4, _ := multiClient.Cluster("").CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	//for _, v := range p4.Items {
	//	fmt.Println(v.Name)
	//}
}
