package public

import (
	"fmt"
	"github.com/filedrive-team/filplus-info/api"
	"github.com/filedrive-team/filplus-info/common"
	"github.com/filedrive-team/filplus-info/errormsg"
	"github.com/filedrive-team/filplus-info/models"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	ClientAllowanceGrantedDailyKey = "client_allowance_granted_daily_%d"
)

func DataCapAllocatedList(c *gin.Context) {
	params := new(api.DataCapAllocatedParam)
	c.BindJSON(params)

	addresses := strings.Split(params.Address, ",")
	res, err := models.ClientList(params.ClientName, params.ClientAddress, addresses, &params.PaginationParams)
	if err != nil {
		api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
	} else {
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), res)
	}
	return
}

func ClientAllowanceGrantedDaily(c *gin.Context) {
	params := new(api.GrantedParam)
	c.BindJSON(params)

	key := fmt.Sprintf(ClientAllowanceGrantedDailyKey, params.Limit)
	if value, ok := common.GlobalCache.Get(key); ok {
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), value)
		return
	}

	res, err := models.GetClientAllowanceGrantedDaily(params.Limit)
	if err != nil {
		api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
	} else {
		common.GlobalCache.Set(key, res, 30*time.Minute)
		api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), res)
	}
	return
}
