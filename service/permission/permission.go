package permission

import (
	"time"
	"wecom-jump-server/config"
	"wecom-jump-server/service/wxhelper"
)

type GetPermissionConfigReply struct {
	NowTime   int64
	NonceStr  string
	Signature string
}

func GetPermissionConfig(url string) (*GetPermissionConfigReply, error) {
	wxHelper := wxhelper.NewWXHelper(config.AppID, config.AgentID, config.Secret)
	ticket, err := wxHelper.GetJsApiTicket("")
	if err != nil {
		return nil, err
	}

	nowTime := time.Now().UnixMilli() / 1000
	nonceStr := "asdfasdfasd"
	signature := wxHelper.GetJSSDKSignature(ticket, nonceStr, url, nowTime)

	reply := &GetPermissionConfigReply{
		NowTime:   nowTime,
		NonceStr:  nonceStr,
		Signature: signature,
	}
	return reply, nil
}
