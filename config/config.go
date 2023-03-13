package config

import "os"

var (
	Domain  string
	AppID   string
	AgentID string
	Secret  string
)

func init() {
	if Domain == "" {
		Domain = os.Getenv("WECOM_JUMP_DOMAIN")
	}

	if AppID == "" {
		AppID = os.Getenv("WECOM_JUMP_APPID")
	}

	if AgentID == "" {
		AgentID = os.Getenv("WECOM_JUMP_AGENTID")
	}

	if Secret == "" {
		Secret = os.Getenv("WECOM_JUMP_SECRET")
	}

	if Domain == "" {
		panic("Domain配置缺失")
	}

	if AppID == "" {
		panic("AppID配置缺失")
	}

	if AgentID == "" {
		panic("AgentID配置缺失")
	}

	if Secret == "" {
		panic("Secret配置缺失")
	}
}
