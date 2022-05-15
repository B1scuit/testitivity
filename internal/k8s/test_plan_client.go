package k8s

import (
	"context"

	"github.com/B1scuit/testitivity/domain"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type TestPlansInterface interface {
	List(context.Context, metav1.ListOptions) (*domain.TestPlanList, error)
	Get(context.Context, string, metav1.GetOptions) (*domain.TestPlan, error)
	Create(context.Context, *domain.TestPlan) (*domain.TestPlan, error)
	Watch(context.Context, metav1.ListOptions) (watch.Interface, error)
}

func (c *Client) TestPlans(ns string) TestPlansInterface {
	return &testPlanClient{
		restClient: c.restClient,
		ns:         ns,
	}
}

type testPlanClient struct {
	restClient rest.Interface
	ns         string
}

func (c *testPlanClient) List(ctx context.Context, opts metav1.ListOptions) (*domain.TestPlanList, error) {
	result := domain.TestPlanList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("testplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *testPlanClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*domain.TestPlan, error) {

	result := domain.TestPlan{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("testplans").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *testPlanClient) Create(ctx context.Context, testplan *domain.TestPlan) (*domain.TestPlan, error) {
	result := domain.TestPlan{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("testplans").
		Body(testplan).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *testPlanClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("testplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}
