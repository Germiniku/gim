/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:43:22
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 11:06:46
 */
package conf

import (
	"log"

	"github.com/pelletier/go-toml"
)

var AppConfig appConfig

type appConfig struct {
	Redis  redisConfig
	Mongo  mongoConfig
	Server serverConfig
}

type serverConfig struct {
	Host string
	Port int
	Mode string
}

type redisConfig struct {
	Addr        string
	Password    string
	DB          int
	DialTimeout int
}

type mongoConfig struct {
	Addr string
}

func Init(path string) (err error) {
	config, err := toml.LoadFile(path)
	if err != nil {
		log.Println("Load Config File Failed,Err:", err.Error())
		return
	}
	err = config.Unmarshal(&AppConfig)
	if err != nil {
		log.Println("Config File Unmarshal AppConfig Failed,Err:", err.Error())
	}
	return
}
