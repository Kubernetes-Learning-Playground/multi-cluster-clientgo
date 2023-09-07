package client

import (
	"context"
	"fmt"
	"golanglearning/new_project/multi_cluster_client/pkg/model"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestRemote(t *testing.T) {
	rn := &model.RemoteNode{
		Host:     "xxx",
		Password: "xxx",
		User:     "root",
		Port:     "22",
	}
	client, _ := GetClientByRemoteKubeConfig(rn, "./config", true)
	p, _ := client.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	for _, v := range p.Items {
		fmt.Println(v.Name)
	}
}
