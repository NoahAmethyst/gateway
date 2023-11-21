package api

import (
	"gateway/common"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	common.SuccessResponse(c, nil)
	return
}
