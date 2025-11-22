//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/server"
	"github.com/go-kratos/kratos-layout/internal/service"
	"github.com/go-kratos/kratos-layout/internal/thirdmodule"
	"github.com/go-kratos/kratos-layout/internal/thirdservice"
	"github.com/go-kratos/kratos/v2"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.ThirdModule, *conf.ThirdService, kratoslog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, thirdmodule.ProviderSet, thirdservice.ProviderSet, service.ProviderSet, newApp))
}
