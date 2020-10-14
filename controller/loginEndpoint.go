/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:22:52
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 17:20:28
 */
package controller

import (
	"gim-route-api/model"
	"gim-route-api/service"
	"gim-route-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginEndpoint(c *gin.Context) {
	var userLoginParams *model.UserLogin
	if err := c.BindJSON(userLoginParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40001,
			"msg":  "请求参数出错",
			"data": nil,
		})
		return
	}
	svc := service.NewUserService()
	user := svc.Login(c, userLoginParams)
	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40002,
			"msg":  "用户不存在",
			"data": nil,
		})
		return
	}
	data := utils.StructToMapFilterField(user,"Password")
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "登录成功",
		"data": data,
	})
	return
}
