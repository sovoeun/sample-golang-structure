package initialize

import (
	"example/sample/global"
	"example/sample/router"
	"github.com/gin-gonic/gin"
)

func AdminRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)

	routers := gin.Default()
	err := routers.SetTrustedProxies([]string{"192.168.1.26"})
	if err != nil {
		return nil
	}
	publicGroup := routers.Group("")
	router.AdminRoute(publicGroup)

	return routers
}
