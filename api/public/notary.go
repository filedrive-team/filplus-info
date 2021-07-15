package public

import (
	"fmt"
	"github.com/filedrive-team/filplus-info/api"
	"github.com/filedrive-team/filplus-info/common"
	"github.com/filedrive-team/filplus-info/errormsg"
	"github.com/filedrive-team/filplus-info/models"
	"github.com/filedrive-team/filplus-info/types"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	NotaryListCacheKey                 = "notary_list_%d_%d"
	ProportionOfAllowanceKey           = "proportion_of_allowance"
	ProportionOfAllowanceByLocationKey = "proportion_of_allowance_by_location"
)

func NotaryList(c *gin.Context) {
	params := new(types.PaginationParams)
	c.BindJSON(params)

	key := fmt.Sprintf(NotaryListCacheKey, params.Page, params.PageSize)
	if value, ok := common.GlobalCache.Get(key); ok {
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), value)
		return
	}
	res, err := models.NotaryList(params)
	if err != nil {
		api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
	} else {
		common.GlobalCache.Set(key, res, 30*time.Minute)
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), res)
	}
	return
}

func ProportionOfAllowance(c *gin.Context) {
	key := ProportionOfAllowanceKey
	if value, ok := common.GlobalCache.Get(key); ok {
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), value)
		return
	}
	res, err := models.GetProportionOfAllowance()
	if err != nil {
		api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
	} else {
		common.GlobalCache.Set(key, res, 30*time.Minute)
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), res)
	}
	return
}

func ProportionOfAllowanceByLocation(c *gin.Context) {
	key := ProportionOfAllowanceByLocationKey
	if value, ok := common.GlobalCache.Get(key); ok {
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), value)
		return
	}
	res, err := models.GetProportionOfAllowanceByLocation()
	if err != nil {
		api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
	} else {
		common.GlobalCache.Set(key, res, 30*time.Minute)
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), res)
	}
	return
}
