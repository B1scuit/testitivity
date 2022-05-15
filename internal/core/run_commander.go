package core

import (
	"context"

	"github.com/B1scuit/testitivity/domain"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) RunCommander() (*domain.TestPlanList, error) {
	return c.k8sClient.TestPlans("default").List(context.TODO(), v1.ListOptions{})
}
