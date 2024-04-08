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
	config := &rest.Config{
		Host: "YOUR-HOST",
	}
	config.BearerToken = "YOUR-TOKEN"
	config.Insecure = true
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list secrets: %v", err)
	}
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
