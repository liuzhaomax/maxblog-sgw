package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-sgw/internal/api"
	"github.com/redis/go-redis/v9"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine  *gin.Engine
	Handler *api.Handler
	Redis   *redis.Client
}
