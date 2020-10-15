/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:22:52
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:48:46
 */
package controller

import (
	"gim-server/common/app"
	"gim-server/common/bcode"
	"gim-server/common/utils"
	"gim-server/model"
	"gim-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginEndpoint(c *gin.Context) {
	var userLoginParams *model.UserLogin
	resp := app.NewResp(c)
	if err := c.BindJSON(userLoginParams); err != nil {
		resp.ServeJSON(http.StatusOK, bcode.INVAILD_PARAMS, nil)
		return
	}
	svc := service.NewUserService()
	user, code := svc.Login(userLoginParams.Useraname, userLoginParams.Password)
	if user == nil {
		resp.ServeJSON(http.StatusOK, code, nil)
		return
	}
	token, err := utils.GenerateToken(userLoginParams.Useraname, userLoginParams.Password)
	if err != nil {
		resp.ServeJSON(http.StatusOK, bcode.ERROR, nil)
		return
	}
	resp.ServeJSON(http.StatusOK, code, map[string]string{
		"username": user.Username,
		"nickname": user.NickName,
		"token":    token,
	})
	return
}
