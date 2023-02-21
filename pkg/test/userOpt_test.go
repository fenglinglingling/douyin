package test

import (
	"douyin/pkg/define"
	"douyin/pkg/helper"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFavoriteOpt(t *testing.T) {
	m := define.M{
		"Token":      "testToken",
		"VideoId":    1,
		"ActionType": 1,
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/favorite/action", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestCommentOpt(t *testing.T) {
	m := define.M{
		"VideoId":     1,
		"ActionType":  1,
		"CommentText": "testCommentText",
		"CommentId":   1,
		"Token":       "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/comment/action", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestFollowOpt(t *testing.T) {
	m := define.M{
		"ActionType": 1,
		"FollowId":   1,
		"Token":      "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/relation/action", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestGetFavoriteList(t *testing.T) {
	m := define.M{
		"UserId": 1,
		"Token":  "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/favorite/list", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestGetCommentList(t *testing.T) {
	m := define.M{
		"VideoId": 1,
		"Token":   "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/comment/list", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestGetFollowList(t *testing.T) {
	m := define.M{
		"UserId": 1,
		"Token":  "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/relation/follow/list", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestGetFollowerList(t *testing.T) {
	m := define.M{
		"UserId": 1,
		"Token":  "testToken",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/douyin/relation/follower/list", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
