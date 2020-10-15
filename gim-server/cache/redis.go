/*
 * @Descripttion:初始化Redis客户端连接
 * @Author: Sun
 * @Date: 2020-10-14 14:32:04
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 15:20:36
 */
package cache

import (
	"gim-server/conf"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func Init() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        conf.AppConfig.Redis.Addr,
		DB:          conf.AppConfig.Redis.DB,
		Password:    conf.AppConfig.Redis.Password,
		DialTimeout: time.Second * time.Duration(conf.AppConfig.Redis.DialTimeout),
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Println("Connect Redis Failed,Err:", err.Error())
		return
	}
	log.Println("Connect to Redis")
	return
}
