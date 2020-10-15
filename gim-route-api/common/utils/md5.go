/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 17:46:22
 * @LastEditors:
 * @LastEditTime: 2020-10-15 17:46:26
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// Md5File md5_file()
func Md5File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}
