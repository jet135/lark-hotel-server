package handlers

import (
	"encoding/json"
	"fmt"
	"lark-hotel-server/internal/lark/initialization"
	"lark-hotel-server/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// SqlAgentAction Invoke the sql ai agent to respond
type SqlAgentAction struct {
}

func (*SqlAgentAction) Execute(a *ActionInfo) bool {
	if sqlAgentKey, found := utils.EitherCutPrefix(a.info.qParsed,
		"sa"); found {
		ctx := *a.ctx
		larkConfig := initialization.GetLarkConfig()
		if larkConfig.SqlAgentApiUrl == "" || larkConfig.SqlAgentApiToken == "" {
			replyMsg(ctx, "agent服务未配置", *a.info.msgId)
			return false
		}
		sqlAgentKey = strings.TrimSpace(sqlAgentKey)
		message := sqlAgentKey
		// 按空格分割
		keys := strings.Split(sqlAgentKey, " ")
		recursionLimit := 10
		if len(keys) == 2 {
			message = keys[0]
			recursionLimit = gconv.Int(keys[1])
		}

		// 构建请求体
		requestBody := g.Map{
			"message":         message,
			"recursion_limit": recursionLimit,
		}

		// 发送POST请求
		client := g.Client()
		client.SetHeader("Content-Type", "application/json")
		client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", larkConfig.SqlAgentApiToken))
		response := client.PostBytes(
			ctx,
			larkConfig.SqlAgentApiUrl,
			requestBody,
		)
		var res map[string]interface{}
		_ = json.Unmarshal(response, &res)
		content := res["content"].(string)
		// 去掉content的换行符
		content = strings.ReplaceAll(content, "\n", "  ")
		replyMsg(ctx, content, *a.info.msgId)
		return false
	}
	return true
}
