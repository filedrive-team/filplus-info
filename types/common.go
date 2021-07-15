package types

import "github.com/shopspring/decimal"

// code for common response type
const (
	SuccessCode   = 0
	ErrorCode     = 1
	ForbiddenCode = 2
	ExpireCode    = 3
)

// CommonResp - api common response type
// success - {"code": 0, "msg": "", "data": {}, "error": null}
// error - {"code": 1, "msg": "", "data": {}, "error": null}
// forbidden - {"code": 2, "msg": "", "data": {}, "error": null}
// expire - {"code": 3, "msg": "", "data": {}, "error": null}
type CommonResp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type CommonList struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type ProportionOfSomething struct {
	Name  string          `json:"name"`
	Value decimal.Decimal `json:"value"`
}
