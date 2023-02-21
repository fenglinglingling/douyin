package test

import (
	"douyin/pkg/define"
	"douyin/pkg/helper"
	"encoding/json"
	"fmt"
	"testing"
)

func TestPublishVideo(t *testing.T) {
	m := define.M{
		"Token": "testToken",
		"Title": "testTitle",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/publish/action", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestGetPublishVideoList(t *testing.T) {
	m := define.M{
		"Token": "testToken",
		"Title": "testTitle",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/publish/list", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestFeedVideoList(t *testing.T) {
	m := define.M{
		"Token":    "testToken",
		"LastTime": 1,
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/publish/", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
