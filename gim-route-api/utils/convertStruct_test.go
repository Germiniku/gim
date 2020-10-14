package utils

import (
	"encoding/json"
	"testing"
)

type User struct {
	Username string
	Password string
}

func TestStructToMapFilterField(t *testing.T) {
	user := User{"username","password"}
	result := StructToMapFilterField(user,"Password","Username")
	currentResult := make(map[string]interface{})
	currentResult["username"] = "username"
	res, err := json.Marshal(&result)
	if err != nil{
		t.Log("json失败")
		t.Failed()
	}
	t.Log(string(res))
	if string(res) == `{}`{
		t.Log("测试成功")
	}
}