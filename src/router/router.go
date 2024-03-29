package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/maxblog-sgw/internal/middleware"
	"github.com/liuzhaomax/maxblog-sgw/src/utils"
)

func Register(root *gin.RouterGroup, mw *middleware.Middleware) {
	root.GET("/login", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))
	root.POST("/login", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))

	root.Use(mw.Auth.ValidateToken())
	root.DELETE("/login", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))
	routerUser := root.Group("/users")
	{
		routerUser.Any("/*url", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))
	}
}
