package spider_api

import (
	"errors"
	"gateway/cluster/spider_svc"
	"gateway/common"
	"gateway/protocol/spider_pb"
	mobile "github.com/floresj/go-contrib-mobile"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// image base region
	random = "random"
	cn     = "cn"
	us     = "us"
	// return type
	redirect = "redirect"
	json     = "json"
)

var r int

func BingWallpaper(c *gin.Context) {
	imageBaseRegion := c.Param("region")
	returnType := c.Param("return_type")

	cli := spider_svc.SvcCli()
	req := &spider_pb.SpiderReq{IsMobile: mobile.GetDevice(c).Mobile()}

	var resp *spider_pb.SpiderResp
	var err error

	switch imageBaseRegion {
	case cn:
		resp, err = cli.GetCNBingWallPaper(c, req)
	case us:
		resp, err = cli.GetUSBingWallPaper(c, req)
	case random:
		if r%2 == 0 {
			resp, err = cli.GetCNBingWallPaper(c, req)
		} else {
			resp, err = cli.GetUSBingWallPaper(c, req)
		}
		r++
	default:
		resp, err = cli.GetUSBingWallPaper(c, req)
	}

	if err != nil {
		common.FailedResponse(c, err)
	} else {
		if returnType == redirect {
			c.Redirect(http.StatusMovedPermanently, resp.Url)
		} else if returnType == json {
			common.SuccessResponse(c, map[string]string{
				"imageUrl": resp.Url,
			})
		} else {
			common.FailedResponse(c, errors.New("unsupported return type"))
		}
	}

	return
}
