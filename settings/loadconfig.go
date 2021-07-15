package settings

import (
	"fmt"
	"github.com/filedrive-team/filplus-info/settings/settingtypes"
	"github.com/filedrive-team/filplus-info/utils"
	"github.com/jinzhu/configor"
)

var AppConfig *settingtypes.AppConfig
var tomlConfig *settingtypes.TomlConfig

func Setup(configFile string) {
	tomlConfig = new(settingtypes.TomlConfig)
	err := configor.Load(tomlConfig, configFile)
	if err != nil {
		panic(fmt.Sprintf("fail to load app config:\n %v\n", err))
	}
	AppConfig = new(settingtypes.AppConfig)
	AppConfig.App = tomlConfig.App
	switch tomlConfig.App.Runmode {
	case "development":
		utils.StructSubCopy(&tomlConfig.Development, &AppConfig.TomlEnv)
	case "test":
		utils.StructSubCopy(&tomlConfig.Test, &AppConfig.TomlEnv)
	case "product":
		utils.StructSubCopy(&tomlConfig.Product, &AppConfig.TomlEnv)
	}
}
