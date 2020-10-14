/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:23:30
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 17:04:49
 */
package controller

import (
	"gim-route-api/model"
	"gim-route-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(c *gin.Context) {
	var userLoginReqParams *model.UserRegister
	svc := service.NewUserService()
	if err := c.BindJSON(userLoginReqParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40001,
			"msg":  "请求参数出错",
			"data": nil,
		})
		return
	}
	ok := svc.Register(c, userLoginReqParams)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "成功",
		"data": nil,
	})
	return
}
