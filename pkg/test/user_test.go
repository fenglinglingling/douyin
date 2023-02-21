package test

import (
	"douyin/pkg/define"
	"douyin/pkg/helper"
	"encoding/json"
	"fmt"
	"testing"
)

var userServiceAddr = "http://127.0.0.1:8080"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "testName",
		"password": "testPassword",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestUserRegister(t *testing.T) {
	m := define.M{
		"username": "testName",
		"password": "testPassword",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/user/register", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
