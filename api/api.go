package api

import (
	"github.com/filedrive-team/filplus-info/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, types.CommonResp{
		Code: types.SuccessCode,
		Msg:  msg,
		Data: data,
	})
}

func JSONError(c *gin.Context, msg string, err interface{}) {
	c.JSON(http.StatusOK, types.CommonResp{
		Code:  types.ErrorCode,
		Msg:   msg,
		Error: err,
	})
}

func JSONExpire(c *gin.Context, msg string, err interface{}) {
	c.JSON(http.StatusOK, types.CommonResp{
		Code:  types.ExpireCode,
		Msg:   msg,
		Error: err,
	})
}

func JSONForbidden(c *gin.Context, msg string, err interface{}) {
	c.JSON(http.StatusOK, types.CommonResp{
		Code:  types.ForbiddenCode,
		Msg:   msg,
		Error: err,
	})
}
