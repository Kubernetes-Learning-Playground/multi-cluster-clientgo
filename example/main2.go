package main

import (
	"context"
	"fmt"
	"golanglearning/new_project/multi_cluster_client/pkg/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	rn := &client.RemoteNode{
		Host: "xxx",
		Password: "xxx",
		User: "root",
		Port: "22",
	}
	c ,_ := client.GetClientByRemoteKubeConfig(rn, "./config", true)
	p , _ := c.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	for _, v := range p.Items {
		fmt.Println(v.Name)
	}
}
