## kubernetes的多集群client-go。
### 项目思路与功能
项目背景：本项目对原生的client-go包进行扩展封装，实现"**多集群**"的client-go。调用方仅需要维护config.yaml配置文件。


### 附注：
1. 目录下创建一个resource文件，把集群的.kube/config文件复制一份放入(记得cluster server需要改成"公网ip")。
2. 本项目支持insecurity模式，所以config文件需要把certificate-authority-data字段删除，否则连接会报错(本身支持tls证书也可以不删除)。(重要！！)
3. 可配置多个.kube/config配置文件。
4. 可读取远程kubeconfig，获取客户端实例

### 配置文件
- **重要** 配置文件可参考config.yaml中配置，调用方只需要关注配置文件中的内容即可。
```yaml
clusters:                     # 集群列表
  - metadata:
      clusterName: cluster1   # 自定义集群名
      insecure: true          # 是否开启跳过tls证书认证
      configPath: /Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config2 # kube config配置文件地址
  - metadata:
      clusterName: cluster2
      insecure: true
      configPath: /Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config1
  - metadata:
      clusterName: cluster3
      insecure: true
      configPath: /Users/zhenyu.jiang/go/src/golanglearning/new_project/multi_cluster_client/resource/config
```

### 使用范例
- 使用配置文件的方式
```go
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
}
```
- 读取远端client实例
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
