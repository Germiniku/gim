/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 10:38:29
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 13:06:52
 */
package utils

import (
	uuid "github.com/satori/go.uuid"
)

// GenerateUniqueId 生成唯一ID
func GenerateUniqueId() string {
	uid := uuid.NewV4()
	return uid.String()
}


