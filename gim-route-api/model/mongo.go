/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 15:11:06
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 16:56:48
 */

package model

import (
	"context"
	"gim-route-api/conf"
	"time"

	mongoDB "go.mongodb.org/mongo-driver/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongo          *mongoDB.Database
	userCollection *mongodb.Collection
)

func Init() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	option := options.Client().ApplyURI(conf.AppConfig.Mongo.Addr)
	option.SetMaxPoolSize(10)
	client, err := mongoDB.Connect(ctx, option)
	if err != nil {
		return
	}
	mongo = client.Database("gim")
	userCollection = mongo.Collection("user")
	return
}
