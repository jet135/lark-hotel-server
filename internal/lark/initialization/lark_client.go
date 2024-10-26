package initialization

import (
	"lark-hotel-server/internal/config"

	lark "github.com/larksuite/oapi-sdk-go/v3"
)

var (
	larkClient    *lark.Client
	todayAppToken string
	todayTableId  string
	larkConfig    config.LarkConfig
)

func LoadLarkClient(config config.LarkConfig) {
	larkConfig = config
	larkClient = lark.NewClient(config.FeishuAppId, config.FeishuAppSecret)
}

func GetLarkClient() *lark.Client {
	return larkClient
}

func GetLarkConfig() config.LarkConfig {
	return larkConfig
}

func LoadTodayAppToken(it string) {
	todayAppToken = it
}

func GetTodayAppToken() string { return todayAppToken }

func LoadTodayTableId(it string) {
	todayTableId = it
}

func GetTodayTableId() string { return todayTableId }
