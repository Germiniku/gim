/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:30:30
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:41:27
 */
package service

import (
	"gim-server/common/bcode"
	"gim-server/model"
)

type UserService interface {
	Register(username, password, nickname string) (int, int)
	Login(username, password string) (user *model.User, code int)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (svc *userService) Register(username, password, nickname string) (int, int) {
	//查询用户是否存在
	if user := model.GetUserByUsername(username); user != nil {
		return 0, bcode.USER_IS_EXSIT
	}
	if id, ok := model.CreateUser(username, password, nickname); !ok {
		return 0, bcode.ERROR
	} else {
		return id, bcode.SUCCESS
	}
}

func (svc *userService) Login(username, password string) (user *model.User, code int) {
	if user = model.GetUserByUsernameAndPassword(username, password); user == nil {
		return nil, bcode.USERNAME_OR_PASSWORD_FAILED
	}
	return user, bcode.SUCCESS
}
