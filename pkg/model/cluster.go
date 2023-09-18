package model

// MetaData 集群对象所需的信息
type MetaData struct {
	// ConfigPath kube config文件
	ConfigPath string `json:"configPath" yaml:"configPath"`
	// Insecure 是否跳过证书认证
	Insecure bool `json:"insecure" yaml:"insecure"`
	// ClusterName 集群名
	ClusterName string `json:"clusterName" yaml:"clusterName"`
	// RemoteMode 远程登入模式
	RemoteMode bool `json:"remoteMode" yaml:"remoteMode"`
	// RemoteNode 远程节点信息
	RemoteNode RemoteNode `json:"remoteNode" yaml:"remoteNode"`
}

type RemoteNode struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Port     string `json:"port" yaml:"port"`
}

// Cluster 集群对象
type Cluster struct {
	MetaData MetaData `json:"metadata" yaml:"metadata"`
}
