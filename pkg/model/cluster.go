package model

// MetaData 集群对象所需的信息
type MetaData struct {
	ConfigPath  string                 `json:"configPath" yaml:"configPath"` // kube config文件
	Insecure    bool                   `json:"insecure" yaml:"insecure"`     // 是否跳过证书认证
	ClusterName string                 `json:"clusterName" yaml:"clusterName"`
}

// Cluster 集群对象
type Cluster struct {
	MetaData MetaData `json:"metadata" yaml:"metadata"`
}
