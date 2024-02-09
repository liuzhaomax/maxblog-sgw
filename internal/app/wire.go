//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-sgw/internal/api"
	"github.com/liuzhaomax/maxblog-sgw/internal/core"
	"github.com/liuzhaomax/maxblog-sgw/internal/middleware"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		core.InitLogrus,
		core.InitGinEngine,
		core.InitRedis,
		core.InitTracer,
		core.InitPrometheusRegistry,
		api.APISet,
		core.LoggerSet,
		middleware.MwsSet,
		middleware.MiddlewareSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
