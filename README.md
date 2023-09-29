## Multi-cluster client-go for kubernetes
<a href="./README.md">English</a> | <a href="./README-zh.md">简体中文</a>
### Introduction
Project background: This project extends the native client-go package to implement the "**multi-cluster**" client-go SDK(clientSet dynamicClient discoveryClient). 

The caller only needs to maintain the config.yaml configuration file.

### P.S.:
1. Create a resource file in the directory, copy the .kube/config file of the cluster and put it in (remember that the cluster server needs to be changed to "public network ip").
2. This project supports insecurity mode, so the certificate-authority-data field needs to be deleted in the config file, otherwise the connection will report an error (it does not need to delete it if it supports TLS certificate). (important!!)
3. Multiple .kube/config configuration files can be configured.
4. Can read remote kubeconfig and obtain client instance

### Configuration file
- **Important** The configuration file can refer to the configuration in config.yaml. The caller only needs to pay attention to the content in the configuration file.
```yaml
clusters:                     # Cluster list
  - metadata:
      clusterName: cluster1   # Custom cluster name
      insecure: true          # Whether to enable skipping tls certificate authentication
      configPath: /Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config2 # kube config配置文件地址
  - metadata:
      clusterName: cluster2
      insecure: true
      configPath: /Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config1
  - metadata:
      clusterName: cluster3
      insecure: true
      configPath: /Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config
  - metadata:
      clusterName: cluster4
      insecure: true
      configPath: ./config
      remoteMode: true        # Remotely read kubeconfig mode
      remoteNode:             # Remote node login information
        host: xxx
        password: xxx
        user: root
        port: 22
```

### Usage examples
- use configuration files
```go
func main() {
    // 1. Create a file
    c, _ := config.BuildConfig("./config.yaml")
    multiClient, _ := client.NewForConfig(c)
    // 2. It is almost the same as the native client-go when used, except that you need to specify the cluster name.
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
}
```
- Read the remote client instance
```go
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
```
