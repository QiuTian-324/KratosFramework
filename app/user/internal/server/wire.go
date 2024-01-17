package server

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"go-Hermes/app/user/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// 创建一个 Consul 默认配置
	c := consulAPI.DefaultConfig()

	// 根据传入的配置覆盖 Consul 配置的一些属性
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	// 使用配置创建 Consul 客户端
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}

	// 使用 Consul 客户端创建 Registrar（服务注册实例）
	r := consul.New(cli, consul.WithHealthCheck(true))

	return r
}
