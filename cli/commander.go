package main

import (
	"path"

	"github.com/B1scuit/testitivity/internal/cli"
	"github.com/B1scuit/testitivity/internal/core"
	"github.com/B1scuit/testitivity/internal/k8s"
	"go.uber.org/zap"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	log, _ := zap.NewDevelopment()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", path.Join(homedir.HomeDir(), ".kube", "config"))
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
		Log:  log,
		Core: core,
	})

	cliClient.Run()
}
