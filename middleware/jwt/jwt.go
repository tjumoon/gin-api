package jwt

import (
	"github.com/gin-gonic/gin"
	"time"
	"gin-api/utils/e"
	"gin-api/utils/jwt"
	"net/http"
	"gin-api/controllers/response"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		token := c.Request.Header.Get("token")
		if token == "" {
			code = e.USER_INVALID_PARAMS
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.USER_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.USER_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			response.ErrorJSON(c, http.StatusUnauthorized, code)
			c.Abort()
			return
		}
		c.Next()
	}
}