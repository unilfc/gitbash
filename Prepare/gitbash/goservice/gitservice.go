package main

import (
    "context"
    "fmt"
    "k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/rest"
)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        panic(err)
    }

    // Create a Kubernetes client
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }

    // Example: List Pods in the "default" namespace
    pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        panic(err)
    }

    for _, pod := range pods.Items {
        fmt.Println("Pod Name:", pod.Name)
    }
}