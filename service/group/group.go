package group

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"wecom-jump-server/config"
	"wecom-jump-server/service/wxhelper"
)

func CreateGroup(groupName string, owner string, userList []string) (string, error) {
	wxHelper := wxhelper.NewWXHelper(config.AppID, config.AgentID, config.Secret)
	accessToken, err := wxHelper.GetAccessToken()
	if err != nil {
		return "", err
	}

	postData := gin.H{
		"name":     groupName,
		"owner":    owner,
		"userList": userList,
	}

	postDataBytes, err := json.Marshal(postData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(wxhelper.BaseUrl+"/cgi-bin/appchat/create?access_token="+accessToken, "application/json", bytes.NewBuffer(postDataBytes))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg := gjson.GetBytes(responseData, "errmsg").String()
		return "", errors.New(errMsg)
	}

	chatID := gjson.GetBytes(responseData, "chatid").String()
	return chatID, nil
}
