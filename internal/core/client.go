package core

import "github.com/B1scuit/testitivity/internal/k8s"

type k8sClientInterface interface {
	TestPlans(string) k8s.TestPlansInterface
}

type ClientOptions struct {
	K8sClient k8sClientInterface
}

type Client struct {
	k8sClient k8sClientInterface
}

func New(opts *ClientOptions) *Client {
	return &Client{
		k8sClient: opts.K8sClient,
	}
}
