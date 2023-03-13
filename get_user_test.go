package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"wecom-jump-server/config"
	"wecom-jump-server/service/wxhelper"
)

func TestGetUser(t *testing.T) {
	appID, agentID, secret := config.AppID, config.AgentID, config.Secret
	helper := wxhelper.NewWXHelper(appID, agentID, secret)
	accessToken, err := helper.GetAccessToken()
	if err != nil {
		t.Log(err)
		return
	}

	postDataBytes := []byte("{\"mobile\": \"xxxxxxxxx\"}")
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/user/getuserid?access_token="+accessToken, "application/json", bytes.NewBuffer(postDataBytes))
	if err != nil {
		t.Log(err)
		return
	}

	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("all is:", string(all))
}

func TestCreateGroup(t *testing.T) {
	appID, agentID, secret := config.AppID, config.AgentID, config.Secret
	helper := wxhelper.NewWXHelper(appID, agentID, secret)
	accessToken, err := helper.GetAccessToken()
	if err != nil {
		t.Log(err)
		return
	}

	postDataBytes := []byte(`{"name": "测试群", "owner":"8239","userlist": ["8239", "6915", "8109"]}`)
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token="+accessToken, "application/json", bytes.NewBuffer(postDataBytes))
	if err != nil {
		t.Log(err)
		return
	}

	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("all is:", string(all))
}

func TestPushMessage(t *testing.T) {
	appID, agentID, secret := config.AppID, config.AgentID, config.Secret
	helper := wxhelper.NewWXHelper(appID, agentID, secret)
	accessToken, err := helper.GetAccessToken()
	if err != nil {
		t.Log(err)
		return
	}

	postDataBytes := []byte(`{"chatid": "xxxxxxx", "msgtype":"text","text": {"content":"xxxxxxxx"}, "safe":0}`)
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token="+accessToken, "application/json", bytes.NewBuffer(postDataBytes))
	if err != nil {
		t.Log(err)
		return
	}

	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("all is:", string(all))
}
