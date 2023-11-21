package constant

import (
	"github.com/gin-gonic/gin"
	"os"
)

const (
	//Gin console model
	GinMode = "GIN_MODE"
)

func GetGinMode() string {
	mode := os.Getenv(GinMode)
	if mode == gin.ReleaseMode || mode == gin.DebugMode || mode == gin.TestMode {
		return mode
	}
	return gin.DebugMode
}
