/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 14:24:54
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 17:10:04
 */
package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	NickName string `bson:"nickname"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
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

func CreateUser(userParams *UserRegister) bool {
	user := newUser(userParams.Username, userParams.Password, userParams.NickName)
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("createUser Failed,Err:", err)
		return false
	}
	return true
}

func newUser(username, password, nickname string) User {
	return User{
		username,
		password,
		nickname,
	}
}
