package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func SuccessResponse(ctx *gin.Context, data any) {
	resp := Resp{Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func FailedResponse(ctx *gin.Context, err error) {
	resp := Resp{Message: err.Error()}
	ctx.JSON(http.StatusInternalServerError, resp)
}
