/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:24:54
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:18:20
 */
package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	NickName string `bson:"nickname"`
}

type UserLogin struct {
	Useraname string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func GetUserByUsername(username string) *User {
	var user *User
	if err := userCollection.FindOne(context.TODO(), bson.M{
		"username": username,
	}).Decode(user); err != nil {
		return nil
	}
	return user
}

func GetUserByUsernameAndPassword(username, password string) *User {
	var user *User
	if err := userCollection.FindOne(context.TODO(), bson.M{
		"username": username,
		"password": password,
	}).Decode(user); err != nil {
		return nil
	}
	return user
}

func CreateUser(username, password, nickname string) (int, bool) {
	user := newUser(username, password, nickname)
	res, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return 0, false
	}
	id := res.InsertedID.(int)
	return id,true
}

func newUser(username, password, nickname string) User {
	return User{
		username,
		password,
		nickname,
	}
}
