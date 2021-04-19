package svc

import (
	"api/internal/config"
	"api/transformer"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),  // 手动代码
	}
}
