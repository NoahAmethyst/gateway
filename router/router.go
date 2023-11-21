package router

import (
	"gateway/api"
	"gateway/api/spider_api"
	"gateway/common"
	mobile "github.com/floresj/go-contrib-mobile"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

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

	_spider := _api.Group("spider")

	{
		_spider.GET("bing/wallpaper/:region/:return_type", spider_api.BingWallpaper)
	}
	return r
}
