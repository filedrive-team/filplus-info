package errormsg

import (
	"github.com/gin-gonic/gin"
)

const (
	SearchSuccess int = iota
	SearchFailed
	SearchError
)
const (
	ZH  string = "zh"
	EN  string = "en"
	JPN string = "jpn"
)

type MultiLang struct {
	ZH  string `json:"zh"`
	EN  string `json:"en"`
	JPN string `json:"jpn"`
}

// ByCtx - give appropriate message according to custom http header "X-Language"
func ByCtx(c *gin.Context, key int) string {
	langcode := c.GetHeader("X-Language")
	return ByLangcode(langcode, key)
}

func ByLangcode(langcode string, key int) string {
	msgData, ok := msgmap[key]
	if !ok {
		return ""
	}
	if langcode == EN {
		return msgData.EN
	}
	if langcode == JPN {
		return msgData.JPN
	}

	return msgData.ZH
}

var msgmap = map[int]MultiLang{
	SearchSuccess: MultiLang{"查询成功", "Query succeeded.", "検索成功"},
	SearchFailed:  MultiLang{"查询失败", "Query failed.", "検索失敗"},
	SearchError:   MultiLang{"查询错误", "Query error.", "検索エラー"},
}
