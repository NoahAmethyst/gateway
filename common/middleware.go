package common

import (
	"gateway/utils/log"
	mobile "github.com/floresj/go-contrib-mobile"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqPlatform := mobile.GetDevice(c).Platform()
		reqIp := c.ClientIP()
		log.Info().Msgf("Request from %s use %s agent [%s]", reqIp, reqPlatform, c.GetHeader("User-Agent"))
		c.Next()
	}
}
