package api

import (
	"github.com/filedrive-team/filplus-info/types"
	"github.com/shopspring/decimal"
)

type DataCapAllocatedParam struct {
	types.PaginationParams
	Address       string `json:"address"`
	ClientName    string `json:"client_name"`
	ClientAddress string `json:"client_address"`
}

type ClientDealParam struct {
	types.PaginationParams
	ClientAddress string          `json:"client_address"`
	StartEpoch    int64           `json:"start_epoch"`
	Allowance     decimal.Decimal `json:"allowance"`
}

type GrantedParam struct {
	Limit int64 `json:"limit"`
}
