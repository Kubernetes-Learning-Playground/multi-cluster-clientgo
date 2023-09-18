package main

import (
	"context"
	"fmt"
	"github.com/practice/multi_cluster_client/pkg/client"
	"github.com/practice/multi_cluster_client/pkg/config"
	"github.com/practice/multi_cluster_client/pkg/model"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {

	// 1. 创建所需的struct

	cfg := config.NewConfig()
	cfg.Clusters = make([]model.Cluster, 0)

	cc := model.Cluster{
		MetaData: model.MetaData{
			ClusterName: "cluster1",
			ConfigPath:  "/Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config",
			Insecure:    true,
		},
	}
	cfg.Clusters = append(cfg.Clusters, cc)

	multiClient, err := client.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 2. 使用时与原生client-go几乎相同，只是需要多指定集群名，
	p1, err := multiClient.Cluster("cluster1").CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range p1.Items {
		fmt.Println(v.Name)
		fmt.Println("aaa")
	}

}
