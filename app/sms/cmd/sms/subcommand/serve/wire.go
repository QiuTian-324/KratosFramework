//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package serve

import (
	"akita/quantum_cat/app/sms/internal/biz"
	"akita/quantum_cat/app/sms/internal/conf"
	"akita/quantum_cat/app/sms/internal/dao"
	"akita/quantum_cat/app/sms/internal/data"
	"akita/quantum_cat/app/sms/internal/server"
	"akita/quantum_cat/app/sms/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
