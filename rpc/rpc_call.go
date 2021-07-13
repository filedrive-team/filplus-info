package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/filedrive-team/filplus-info/settings"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func RpcCall(rpcMetchod string, rpcParams interface{}) (result []byte, requetRpcId int, err error) {
	conf := settings.AppConfig
	url := conf.App.FilscanApi + "/rpc/v1"

	logger.Info(url, ", method=", rpcMetchod, ", rpcParams=", rpcParams)
	rpcId := rand.Int()

	rpcJsonMsg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  rpcMetchod,
		"params":  rpcParams,
		"id":      rpcId,
	}
	requetRpcId = rpcId

	mJson, err := json.Marshal(rpcJsonMsg)
	if err != nil {
		logger.Errorf("RpcCall marshal param failed: %v", err)
		return
	}
	contentReader := bytes.NewReader(mJson)
	request, err := http.NewRequest("POST", url, contentReader)
	if err != nil {
		logger.Errorf("RpcCall http request failed: %v", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Errorf("RpcCall client do failed: %v", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("RpcCall ioutil failed: %v", err)
		return
	}

	return body, requetRpcId, nil
}

func checkResponse(responseError []byte) (err error) {
	if responseError != nil {
		type RpcError struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}

		var rpcError RpcError
		json.Unmarshal(responseError, &rpcError)

		logger.Error("rpc error : ", rpcError)
		return errors.New(string(responseError))
	}

	return nil
}
