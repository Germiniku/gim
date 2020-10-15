/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:18:15
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 15:22:11
 */
package router

import (
	"gim-server/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	router.POST("/user/register", controller.RegisterEndpoint)
	router.POST("/user/login", controller.LoginEndpoint)
	return router 
}
