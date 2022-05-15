package k8s

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ClientOptions struct {
	RestConfig *rest.Config
}

type Client struct {
	restClient *rest.RESTClient
}

func New(opts *ClientOptions) *Client {

	AddToScheme(scheme.Scheme)

	crdConfig := *opts.RestConfig
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: groupName, Version: groupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	exampleRestClient, err := rest.UnversionedRESTClientFor(&crdConfig)
	if err != nil {
		panic(err.Error())
	}

	return &Client{
		restClient: exampleRestClient,
	}
}
