package main

import (
	"github.com/B1scuit/testitivity/internal/cli"
	"github.com/B1scuit/testitivity/internal/core"
	"github.com/B1scuit/testitivity/internal/k8s"
	"go.uber.org/zap"

	"k8s.io/client-go/rest"
)

func main() {
	log, _ := zap.NewDevelopment()
	// creates the in-cluster config

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// Init the k8s package
	k8sClient := k8s.New(&k8s.ClientOptions{
		RestConfig: config,
	})

	// create the core package
	core := core.New(&core.ClientOptions{
		K8sClient: k8sClient,
	})

	cliClient := cli.New(&cli.ClientOptions{
		Core: core,
		Log:  log,
	})

	cliClient.Run()
}
