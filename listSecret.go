package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 创建空白配置
	config := &rest.Config{
		Host: "YOUR-HOST",
	}
	config.BearerToken = "YOUR-TOKEN"
	config.Insecure = true
	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// 获取 Secret 列表
	secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list secrets: %v", err)
	}

	// 打印 Secret 信息
	fmt.Println("Secrets:")
	for _, secret := range secrets.Items {
		fmt.Printf("Name: %s\n", secret.Name)
		fmt.Printf("Namespace: %s\n", secret.Namespace)
		fmt.Println("Data:")
		for key := range secret.Data {
			fmt.Printf("- %s: %s\n", key, string(secret.Data[key]))
		}
		fmt.Println("-----")
	}
	fmt.Println("-----")

}
