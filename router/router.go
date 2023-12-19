package router

import (
	"gateway/api"
	"gateway/api/spider_api"
	"gateway/common"
	"gateway/constant"
	"gateway/service"
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

	// Sandbox api

	_test := _api.Group("sandbox")
	{
		_test.GET("stream", service.Stream)
	}
	// Spider api

	_spider := _api.Group("spider")
	{
		_spider.GET("bing/wallpaper/:region/:return_type", spider_api.BingWallpaper)
		_spider.GET("/hot/weibo", spider_api.WeiboHot)
		_spider.GET("/hot/zhihu", spider_api.ZhiHuHot)
		_spider.GET("/hot/36kr", spider_api.D36KrHot)
		_spider.GET("/news/wallstreet", spider_api.WallStreetNews)

	}

	return r
}
