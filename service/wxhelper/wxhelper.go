package wxhelper

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

var BaseUrl = "https://qyapi.weixin.qq.com"

var ticketMap sync.Map

func init() {
	ticketMap = sync.Map{}
}

type WXHelper struct {
	AppID       string
	AgentID     string
	Secret      string
	getTokenUrl string
}

func NewWXHelper(appID, agentID, secret string) *WXHelper {
	wxHelper := new(WXHelper)
	wxHelper.AppID = appID
	wxHelper.AgentID = agentID
	wxHelper.Secret = secret
	wxHelper.getTokenUrl = BaseUrl + "/cgi-bin/gettoken?corpid=" + appID + "&corpsecret=" + secret
	return wxHelper
}

func (h *WXHelper) GetAccessToken() (string, error) {
	resp, err := http.Get(h.getTokenUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("respBody is :", string(respBody))
	accessToken := gjson.Get(string(respBody), "access_token").String()
	return accessToken, nil
}

func (h *WXHelper) GetJsApiTicket(t string) (string, error) {
	accessToken, err := h.GetAccessToken()
	if err != nil {
		return "", err
	}
	key := accessToken
	if len(accessToken) > 0 {
		if t == "agent_config" {
			key = t + "_" + accessToken
		}

		value, ok := ticketMap.Load(key)
		if ok {
			ticket := value.(Ticket)
			nowTime := time.Now().Unix()
			expiresIn := ticket.getExpiresIn()
			if expiresIn-nowTime > 0 {
				return ticket.getTicket(), nil
			}
		}

		ticket, err := h.getJsApiTicketFromWeChatPlatform(accessToken, t)
		if err != nil {
			return "", err
		}

		if ticket != nil {
			ticketMap.Store(ticket.getTicket(), ticket)
			return ticket.getTicket(), nil
		}
	}

	return "", nil
}

func (h *WXHelper) getJsApiTicketFromWeChatPlatform(accessToken, t string) (*Ticket, error) {
	var url string
	if t == "agent_config" {
		url = BaseUrl + "/cgi-bin/ticket/get?access_token=" + accessToken + "&type=" + t
	} else {
		url = BaseUrl + "/cgi-bin/get_jsapi_ticket?access_token=" + accessToken
	}

	now := time.Now().UnixMilli()
	if accessToken != "" {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		fmt.Println("respBody is:", string(respBody))
		if len(respBody) > 0 {
			errCode := gjson.Get(string(respBody), "errcode").Int()
			if errCode == 0 {
				ticketValue := gjson.Get(string(respBody), "ticket").String()
				expiresIn := gjson.Get(string(respBody), "expires_in").Int()
				ticket := NewTicket(ticketValue, now+expiresIn*1000)
				return &ticket, nil
			}
		}
	}

	return nil, nil
}

func (h *WXHelper) GetJSSDKSignature(ticket, nonceStr, url string, timestamp int64) string {
	unEncryptStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, nonceStr, timestamp, url)
	data := []byte(unEncryptStr)
	hash := sha1.Sum(data)
	hashBytes := hash[:]
	encryptStr := hex.EncodeToString(hashBytes)
	return encryptStr
}
