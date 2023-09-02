package client

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestRemote(t *testing.T) {
	rn := &RemoteNode{
		Host: "xxx",
		Password: "xxx",
		User: "root",
		Port: "22",
	}
	client ,_ := GetClientByRemoteKubeConfig(rn, "./config", true)
	p , _ := client.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	for _, v := range p.Items {
		fmt.Println(v.Name)
	}
}
