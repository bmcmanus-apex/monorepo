package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig string

	// If home directory is not empty, set kubeconfig to the path to the kubeconfig file.
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// build configuration to connect to a K8s cluster based on CLI flags and kubeconfig path
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("Falling back to in-cluster config")

		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Failed to create config: %v", err)
		}
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create dynamic client: %v", err)
	}

	// Define CRD
	myCustomResource := schema.GroupVersionResource{
		Group:    "myk8s.io",
		Version:  "v1",
		Resource: "mycustomresources",
	}

	// Setup Informer that watches for changes to resource & maintains a local cache of all resources of this type
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return dynamicClient.Resource(myCustomResource).Namespace("").List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return dynamicClient.Resource(myCustomResource).Namespace("").Watch(context.TODO(), options)
			},
		},
		&unstructured.Unstructured{},
		// resync period of zero means the informer will not resync unless explicitly triggered
		0,
		cache.Indexers{},
	)

	// register callback for different types of events relate to resource
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("New resource added: ", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Resource updated: ", newObj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Resource deleted: ", obj)
		},
	})

	// Start informer
	stop := make(chan struct{})
	defer close(stop)

	// start informer in separate goroutine, allowing it to begin processing events
	go informer.Run(stop)

	// block until informer's local cache is synced with the K8s server
	if !cache.WaitForCacheSync(stop, informer.HasSynced) {
		log.Fatalf("Failed to sync cache")
	}

	fmt.Println("Informer synced and ready")

	<-stop
}
