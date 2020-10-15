/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 17:27:59
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:34:54
 */
package app

import (
	"gim-server/common/bcode"

	"github.com/gin-gonic/gin"
)

type Resp interface {
	ServeJSON(httpCode, errCode int, data interface{})
}

type resp struct {
	c *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(c *gin.Context) Resp {
	return &resp{
		c: c,
	}
}

func (r *resp) ServeJSON(httpCode, errCode int, data interface{}) {
	r.c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  bcode.GetMsg(errCode),
		Data: data,
	})
}
