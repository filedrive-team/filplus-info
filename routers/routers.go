package routers

import (
	"fmt"
	"github.com/filedrive-team/filplus-info/api"
	"github.com/filedrive-team/filplus-info/api/public"
	"github.com/filedrive-team/filplus-info/errormsg"
	"github.com/filedrive-team/filplus-info/middleware/cors"
	"github.com/filedrive-team/filplus-info/settings"

	"github.com/gin-gonic/gin"
)

// InitRouter - initialize routing information
func InitRouter() *gin.Engine {
	conf := settings.AppConfig
	if conf.App.Runmode == settings.RunmodeProd {
		fmt.Printf("app runmode: %s %s", conf.App.Runmode, settings.RunmodeProd)
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.AddCorsHeaders())

	baseGroup := r.Group("/api/v1")
	{
		baseGroup.POST("/version", func(c *gin.Context) {
			msg := errormsg.ByCtx(c, errormsg.SearchSuccess)
			data := settings.Version
			api.JSON(c, msg, data)
		})

		baseGroup.POST("/notaries", public.NotaryList)
		baseGroup.POST("/allocated", public.DataCapAllocatedList)
		baseGroup.POST("/deals", public.ClientDataCapDealList)

		baseGroup.POST("/granted-daily", public.ClientAllowanceGrantedDaily)
		baseGroup.POST("/proportion-of-allowance", public.ProportionOfAllowance)
		baseGroup.POST("/proportion-of-allowance-by-location", public.ProportionOfAllowanceByLocation)

		//adminGroup := baseGroup.Group("/admin")
		//adminGroup.Use(session.CheckAdminSession())
		//{
		//	adminGroup.POST("/ronglaiTransPreUserList", admin.RonglaiTransPreUserList)
		//
		//
		//}

	}

	return r
}
