package handlers

import (
	"context"
	"lark-hotel-server/internal/config"
	"lark-hotel-server/internal/lark/initialization"
	"testing"
)

func init() {
	larkConfig := config.LarkConfig{FeishuAppId: "", FeishuAppSecret: ""}
	initialization.LoadLarkClient(larkConfig)
}

var receiveId = ""

func TestHelpCardTemplate(t *testing.T) {
	_ = MessagesCard(context.Background(), receiveId, HelpCardTemplate())
}

func TestCustomerCardTemplate(t *testing.T) {
	// _ = MessagesCard(context.Background(), receiveId, CustomerInfoCardTemplate(nil))
}
