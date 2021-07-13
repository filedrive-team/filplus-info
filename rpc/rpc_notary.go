package rpc

import (
	"encoding/json"
	"errors"
	"github.com/filedrive-team/filplus-info/models"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

type NotaryDataCap struct {
	TxnId          int64           `pg:"txn_id" json:"txn_id"`
	Address        string          `pg:"address" json:"address"`
	Allowance      decimal.Decimal `pg:"allowance" json:"allowance"`
	GrantAllowance decimal.Decimal `pg:"grant_allowance" json:"grant_allowance"`
}

type NotaryDataCapAllocated struct {
	Epoch     int64           `pg:"epoch" json:"epoch"`
	BlockTime int64           `pg:"block_time" json:"block_time"`
	Address   string          `pg:"address" json:"address"`
	Client    string          `pg:"client" json:"client"`
	Allowance decimal.Decimal `pg:"allowance" json:"allowance"`
}

type MarketDeal struct {
	//tableName            string          `pg:"lotus_market_deal,discard_unknown_columns"`
	Id                   int64           `pg:"id" json:"-"`
	Epoch                int64           `pg:"epoch" json:"epoch"`
	Label                string          `pg:"label" json:"label"`
	Cid                  string          `pg:"cid" json:"cid"`
	DealID               int64           `pg:"dealid" json:"dealid"`
	Client               string          `pg:"client" json:"client"`
	StartEpoch           int64           `pg:"start_epoch" json:"start_epoch"`
	EndEpoch             int64           `pg:"end_epoch" json:"end_epoch"`
	PieceCID             string          `pg:"piece_cid" json:"piece_cid"`
	Provider             string          `pg:"provider" json:"provider"`
	PieceSize            decimal.Decimal `pg:"piece_size" json:"piece_size"`
	VerifiedDeal         bool            `pg:"verified_deal" json:"verified_deal"`
	ClientCollateral     decimal.Decimal `pg:"client_collateral" json:"client_collateral"`
	ProviderCollateral   decimal.Decimal `pg:"provider_collateral" json:"provider_collateral"`
	StoragePricePerEpoch decimal.Decimal `pg:"storage_price_per_epoch" json:"storage_price_per_epoch"`
	BlockTime            int             `pg:"block_time" json:"block_time"`
	ServiceStart         int             `pg:"service_start" json:"service_start_time"`
	ServiceEnd           int             `pg:"service_end" json:"service_end_time"`
}

type NotaryDataCapList struct {
	Total int              `json:"total"`
	List  []*NotaryDataCap `json:"list"`
}

type NotaryDataCapAllocatedList struct {
	Total int                       `json:"total"`
	List  []*NotaryDataCapAllocated `json:"list"`
}

type ClientDataCapDealList struct {
	Total int           `json:"total"`
	List  []*MarketDeal `json:"list"`
}

type ClientDataCapDealNum struct {
	Total int `json:"total"`
}

func GetNotaryList(start, count int) (data *NotaryDataCapList, err error) {
	rpcParams := [...]interface{}{
		start,
		count,
	}
	body, _, err := RpcCall("filscan.NotaryDataCapList", rpcParams)
	if err != nil {
		return
	}

	type jsonRpcResponse struct {
		JsonRpc string             `json:"jsonrpc"`
		Result  *NotaryDataCapList `json:"result"`
		Error   []byte             `json:"error"`
	}

	var jsonResp jsonRpcResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Errorf("rpc unmarshal json failed: %v", err)
		return
	}

	if jsonResp.Result == nil {
		err = errors.New("rpc response.Result is nil ")
		return
	}

	err = checkResponse(jsonResp.Error)
	if err != nil {
		return
	}

	return jsonResp.Result, nil
}

func GetNotaryDataCapAllocatedList(address string, client []string, start, count int) (data *NotaryDataCapAllocatedList, err error) {
	rpcParams := [...]interface{}{
		address,
		client,
		start,
		count,
	}
	body, _, err := RpcCall("filscan.NotaryDataCapAllocatedList", rpcParams)
	if err != nil {
		return
	}

	type jsonRpcResponse struct {
		JsonRpc string                      `json:"jsonrpc"`
		Result  *NotaryDataCapAllocatedList `json:"result"`
		Id      int                         `json:"id"`
		Error   []byte                      `json:"error"`
	}

	var jsonResp jsonRpcResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Errorf("rpc unmarshal json failed: %v", err)
		return
	}

	if jsonResp.Result == nil {
		err = errors.New("rpc response.Result is nil ")
		return
	}

	err = checkResponse(jsonResp.Error)
	if err != nil {
		return
	}

	return jsonResp.Result, nil
}

func GetClientDataCapDealList(client string, startEpoch int64, start, count int) (data *ClientDataCapDealList, err error) {
	rpcParams := [...]interface{}{
		client,
		startEpoch,
		start,
		count,
	}
	body, _, err := RpcCall("filscan.ClientDataCapDealList", rpcParams)
	if err != nil {
		return
	}

	type jsonRpcResponse struct {
		JsonRpc string                 `json:"jsonrpc"`
		Result  *ClientDataCapDealList `json:"result"`
		Id      int                    `json:"id"`
		Error   []byte                 `json:"error"`
	}

	var jsonResp jsonRpcResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Errorf("rpc unmarshal json failed: %v", err)
		return
	}

	if jsonResp.Result == nil {
		err = errors.New("rpc response.Result is nil ")
		return
	}

	err = checkResponse(jsonResp.Error)
	if err != nil {
		return
	}

	return jsonResp.Result, nil
}

func CalcClientDataCapDealNum(client string, startEpoch int64, allowance decimal.Decimal) (data *ClientDataCapDealNum, err error) {
	rpcParams := [...]interface{}{
		client,
		startEpoch,
		allowance,
	}
	body, _, err := RpcCall("filscan.CalcClientDataCapDealNum", rpcParams)
	if err != nil {
		return
	}

	type jsonRpcResponse struct {
		JsonRpc string                `json:"jsonrpc"`
		Result  *ClientDataCapDealNum `json:"result"`
		Id      int                   `json:"id"`
		Error   []byte                `json:"error"`
	}

	var jsonResp jsonRpcResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Errorf("rpc unmarshal json failed: %v", err)
		return
	}

	if jsonResp.Result == nil {
		err = errors.New("rpc response.Result is nil ")
		return
	}

	err = checkResponse(jsonResp.Error)
	if err != nil {
		return
	}

	return jsonResp.Result, nil
}

func NotaryAllowanceList(startEpoch, endEpoch uint64) (data []*models.NotaryAllowance, err error) {
	rpcParams := [...]interface{}{
		startEpoch,
		endEpoch,
	}
	body, _, err := RpcCall("filscan.NotaryAllowanceList", rpcParams)
	if err != nil {
		return
	}

	type jsonRpcResponse struct {
		JsonRpc string                    `json:"jsonrpc"`
		Result  []*models.NotaryAllowance `json:"result"`
		Id      int                       `json:"id"`
		Error   []byte                    `json:"error"`
	}

	var jsonResp jsonRpcResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Errorf("rpc unmarshal json failed: %v", err)
		return
	}

	if jsonResp.Result == nil {
		err = errors.New("rpc response.Result is nil ")
		return
	}

	err = checkResponse(jsonResp.Error)
	if err != nil {
		return
	}

	return jsonResp.Result, nil
}

func ClientAllowanceList(startBlockTime, endBlockTime int64) (data []*models.ClientAllowance, err error) {
	rpcParams := [...]interface{}{
		startBlockTime,
		endBlockTime,
	}
	body, _, err := RpcCall("filscan.ClientAllowanceList", rpcParams)
	if err != nil {
		return
	}

	type jsonRpcResponse struct {
		JsonRpc string                    `json:"jsonrpc"`
		Result  []*models.ClientAllowance `json:"result"`
		Id      int                       `json:"id"`
		Error   []byte                    `json:"error"`
	}

	var jsonResp jsonRpcResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Errorf("rpc unmarshal json failed: %v", err)
		return
	}

	if jsonResp.Result == nil {
		err = errors.New("rpc response.Result is nil ")
		return
	}

	err = checkResponse(jsonResp.Error)
	if err != nil {
		return
	}

	return jsonResp.Result, nil
}
