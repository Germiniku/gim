/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 17:51:24
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:57:20
 */
package auth

import (
	"gim-server/common/app"
	"gim-server/common/bcode"
	"gim-server/common/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := app.NewResp(c)
		token := c.Query("token")
		if token == "" {
			resp.ServeJSON(http.StatusOK, bcode.ERROR_LOST_TOKEN, nil)
			c.Abort()
			return
		}
		_, err := utils.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				resp.ServeJSON(http.StatusOK, bcode.ERROR_AUTH_CHECK_EXPIRED, nil)
			default:
				resp.ServeJSON(http.StatusOK, bcode.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			}
			c.Abort()
			return
		}
	}
}
