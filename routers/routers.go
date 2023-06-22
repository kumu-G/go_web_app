package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app2/logger"
	"web_app2/settings"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(context *gin.Context) {
		context.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
