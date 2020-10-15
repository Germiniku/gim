/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:23:30
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:42:06
 */
package controller

import (
	"gim-server/common/app"
	"gim-server/common/bcode"
	"gim-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
}

func RegisterEndpoint(c *gin.Context) {
	var (
		userLoginReqParams *UserRegister
	)
	resp := app.NewResp(c)
	svc := service.NewUserService()
	if err := c.BindJSON(userLoginReqParams); err != nil {
		resp.ServeJSON(http.StatusOK, bcode.INVAILD_PARAMS, nil)
		return
	}
	id, code := svc.Register(userLoginReqParams.Username, userLoginReqParams.Password, userLoginReqParams.NickName)
	if code != bcode.SUCCESS {
		resp.ServeJSON(http.StatusOK, code, nil)
		return
	}
	resp.ServeJSON(http.StatusOK, code, map[string]int{"id": id})
	return
}
