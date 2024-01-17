//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package serve

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-Hermes/app/user/internal/biz"
	"go-Hermes/app/user/internal/conf"
	"go-Hermes/app/user/internal/dao"
	"go-Hermes/app/user/internal/data"
	"go-Hermes/app/user/internal/server"
	"go-Hermes/app/user/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
