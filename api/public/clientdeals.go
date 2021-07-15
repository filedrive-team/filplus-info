package public

import (
	"github.com/filedrive-team/filplus-info/api"
	"github.com/filedrive-team/filplus-info/errormsg"
	"github.com/filedrive-team/filplus-info/rpc"
	"github.com/filedrive-team/filplus-info/settings"
	"github.com/filedrive-team/filplus-info/types"
	"github.com/filedrive-team/filplus-info/utils"
	"github.com/gin-gonic/gin"
)

func ClientDataCapDealList(c *gin.Context) {
	params := new(api.ClientDealParam)
	c.BindJSON(params)

	offset, size := utils.PaginationHelper(params.Page, params.PageSize, settings.DefaultPageSize)
	total, err := rpc.CalcClientDataCapDealNum(params.ClientAddress, params.StartEpoch, params.Allowance)
	if err != nil {
		api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
		return
	}
	res := &types.CommonList{
		Total: total.Total,
	}
	if total.Total > 0 {
		data, err := rpc.GetClientDataCapDealList(params.ClientAddress, params.StartEpoch, offset, size)
		if err != nil {
			api.JSONError(c, errormsg.ByCtx(c, errormsg.SearchFailed), err.Error())
			return
		}
		res.List = data.List
	}
	api.JSON(c, errormsg.ByCtx(c, errormsg.SearchSuccess), res)
	return
}
