/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:30:30
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 17:18:44
 */
package service

import (
	"gim-route-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(c *gin.Context, params *model.UserRegister) bool
	Login(c *gin.Context, params *model.UserLogin) (user *model.User)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (svc *userService) Register(c *gin.Context, params *model.UserRegister) bool {
	//查询用户是否存在
	if user := model.GetUserByUsername(params.Username); user != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40002,
			"msg":  "用户已经存在",
			"data": nil,
		})
		return false
	}
	model.CreateUser(params)
	return true
}

func (svc *userService) Login(c *gin.Context, params *model.UserLogin) (user *model.User) {
	user = model.GetUserByUsernameAndPassword(params.Useraname, params.Password)
	return user
}
