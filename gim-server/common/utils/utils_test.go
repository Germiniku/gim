/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 10:38:29
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 11:04:15
 */

package utils

import (
	"testing"
)

func TestGenerateUniqueId(t *testing.T) {
	got := GenerateUniqueId()
	t.Log(got)
}
