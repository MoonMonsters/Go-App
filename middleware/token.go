package middleware

import (
	"Go-App/pkg/e"
	"Go-App/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func TokenVer() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("authorization")

		if token == "" {
			util.ResponseWithJson(e.ERROR_AUTH_TOKEN, "", c)
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				util.ResponseWithJson(e.ERROR_AUTH_CHECK_TOKEN_FAIL, "", c)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				util.ResponseWithJson(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, "", c)
				c.Abort()
				return
			} else {
				c.Set("ID", claims.ID)
				c.Next()
			}
		}
	}
}
