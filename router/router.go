package router

import (
	"gateway/api"
	"gateway/api/spider_api"
	"gateway/common"
	"gateway/constant"
	mobile "github.com/floresj/go-contrib-mobile"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	// Setup gin console color
	gin.ForceConsoleColor()

	// Setup gin console format
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Info().Msgf("")
	//}

	r := gin.Default()

	gin.SetMode(constant.GetGinMode())

	// Use mobile resolver to get device
	r.Use(mobile.Resolver())
	// Use logger
	r.Use(common.Logger())
	prefix := "/gateway"

	_api := r.Group(prefix)

	{
		_api.GET("healthcheck", api.HealthCheck)
		_api.POST("healthcheck", api.HealthCheck)
	}

	// Spider api

	_spider := _api.Group("spider")
	{
		_spider.GET("bing/wallpaper/:region/:return_type", spider_api.BingWallpaper)
	}

	return r
}
