package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCorsHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,X-Language")
		if http.MethodOptions == c.Request.Method {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{})
		} else {
			c.Next()
		}
	}
}
