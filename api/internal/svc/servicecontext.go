package svc

import (
	"github.com/lifezq/shorturl/api/internal/config"
	"github.com/lifezq/shorturl/rpc/transform/transformer"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
