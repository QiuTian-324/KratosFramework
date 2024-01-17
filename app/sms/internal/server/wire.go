package server

import (
	"akita/quantum_cat/app/sms/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet 定义了一组服务提供者，用于依赖注入。
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar, NewDiscovery)

// NewRegistrar 创建并返回一个 registry.Registrar 接口的实例，用于服务注册。
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// 从配置中获取 Consul 的相关配置信息
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	// 创建一个 Consul 客户端 cli
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}

	// 使用 Consul 客户端创建一个服务注册实例 r，包括健康检查
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}

// NewDiscovery 创建并返回一个 registry.Discovery 接口的实例，用于服务发现。
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	// 从配置中获取 Consul 的相关配置信息
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Address

	// 创建一个 Consul 客户端 cli
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}

	// 使用 Consul 客户端创建一个服务发现实例 r，包括健康检查
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}
