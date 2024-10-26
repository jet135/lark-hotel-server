package handlers

import (
	"context"
	"encoding/json"
	"lark-hotel-server/internal/utils"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type MsgInfo struct {
	handlerType HandlerType
	msgType     string
	msgId       *string
	chatId      *string
	qParsed     string
	fileKey     string
	imageKey    string
	sessionId   *string
	mention     []*larkim.MentionEvent
}

func (m MsgInfo) MarshalJSON() ([]byte, error) {
	type Alias MsgInfo
	return json.Marshal(&struct {
		HandlerType HandlerType            `json:"handler_type"`
		MsgType     string                 `json:"msg_type"`
		MsgId       *string                `json:"msg_id,omitempty"`
		ChatId      *string                `json:"chat_id,omitempty"`
		QParsed     string                 `json:"q_parsed"`
		FileKey     string                 `json:"file_key"`
		ImageKey    string                 `json:"image_key"`
		SessionId   *string                `json:"session_id,omitempty"`
		Mention     []*larkim.MentionEvent `json:"mention,omitempty"`
	}{
		HandlerType: m.handlerType,
		MsgType:     m.msgType,
		MsgId:       m.msgId,
		ChatId:      m.chatId,
		QParsed:     m.qParsed,
		FileKey:     m.fileKey,
		ImageKey:    m.imageKey,
		SessionId:   m.sessionId,
		Mention:     m.mention,
	})
}

type ActionInfo struct {
	handler *MessageHandler
	ctx     *context.Context
	info    *MsgInfo
}

type Action interface {
	Execute(a *ActionInfo) bool
}

type ProcessedUniqueAction struct { //消息唯一性
}

func (*ProcessedUniqueAction) Execute(a *ActionInfo) bool {
	if a.handler.msgCache.IfProcessed(*a.info.msgId) {
		return false
	}
	a.handler.msgCache.TagProcessed(*a.info.msgId)
	return true
}

type ProcessMentionAction struct { //是否机器人应该处理
}

func (*ProcessMentionAction) Execute(a *ActionInfo) bool {
	// 私聊直接过
	if a.info.handlerType == UserHandler {
		return true
	}
	// 群聊判断是否提到机器人
	if a.info.handlerType == GroupHandler {
		if a.handler.judgeIfMentionMe(a.info.mention) {
			return true
		}
		return false
	}
	return false
}

type EmptyAction struct { /*空消息*/
}

func (*EmptyAction) Execute(a *ActionInfo) bool {
	if len(a.info.qParsed) == 0 {
		sendHelpCard(*a.ctx, a.info.msgId)
		return false
	}
	return true
}

type HelpAction struct { /*帮助*/
}

func (*HelpAction) Execute(a *ActionInfo) bool {
	if _, foundHelp := utils.EitherTrimEqual(a.info.qParsed, "/help",
		"帮助"); foundHelp {
		sendHelpCard(*a.ctx, a.info.msgId)
		return false
	}
	return true
}
