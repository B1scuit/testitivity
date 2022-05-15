package cli

import (
	"github.com/B1scuit/testitivity/domain"
	"go.uber.org/zap"
)

type coreClientInterface interface {
	RunCommander() (*domain.TestPlanList, error)
}

type ClientOptions struct {
	Log *zap.Logger

	Core coreClientInterface
}

type Client struct {
	log  *zap.Logger
	core coreClientInterface
}

func New(opts *ClientOptions) *Client {
	return &Client{
		log: opts.Log,

		core: opts.Core,
	}
}
