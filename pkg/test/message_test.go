package test

import (
	"douyin/pkg/define"
	"douyin/pkg/helper"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	m := define.M{
		"username":   "testName",
		"password":   "testPassword",
		"Token":      "testToken",
		"ToUserId":   1,
		"ActionType": 1,
		"Content":    "testContent",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/message/action", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestMessagelist(t *testing.T) {
	m := define.M{
		"UserId": "testUserId",
		"Token":  "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/message/chat", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
