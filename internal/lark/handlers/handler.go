package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"lark-hotel-server/internal/config"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/service"
	"strings"

	"github.com/google/uuid"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 责任链
func chain(data *ActionInfo, actions ...Action) bool {
	for _, v := range actions {
		if !v.Execute(data) {
			return false
		}
	}
	return true
}

type MessageHandler struct {
	//sessionCache services.SessionServiceCacheInterface
	msgCache MsgCacheInterface
	config   config.LarkConfig
}

func (m MessageHandler) cardHandler(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
	return nil, nil
}

func judgeMsgType(event *larkim.P2MessageReceiveV1) (string, error) {
	msgType := event.Event.Message.MessageType

	switch *msgType {
	case "text", "image", "audio":
		return *msgType, nil
	default:
		return "", fmt.Errorf("unknown message type: %v", *msgType)
	}
}

func (m MessageHandler) msgReceivedHandler(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	handlerType := judgeChatType(event)
	if handlerType == "otherChat" {
		fmt.Println("unknown chat type")
		return nil
	}

	msgType, err := judgeMsgType(event)
	if err != nil {
		fmt.Printf("error getting message type: %v\n", err)
		return nil
	}

	content := event.Event.Message.Content
	msgId := event.Event.Message.MessageId
	rootId := event.Event.Message.RootId
	chatId := event.Event.Message.ChatId
	mention := event.Event.Message.Mentions

	sessionId := rootId
	if sessionId == nil || *sessionId == "" {
		sessionId = msgId
	}
	msgInfo := MsgInfo{
		handlerType: handlerType,
		msgType:     msgType,
		msgId:       msgId,
		chatId:      chatId,
		qParsed:     strings.Trim(parseContent(*content), " "),
		fileKey:     parseFileKey(*content),
		imageKey:    parseImageKey(*content),
		sessionId:   sessionId,
		mention:     mention,
	}
	data := &ActionInfo{
		ctx:     &ctx,
		handler: &m,
		info:    &msgInfo,
	}

	// 记lark_event_operation_log
	var msgInfoJson string
	if jsonContent, err := json.Marshal(msgInfo); err == nil {
		msgInfoJson = string(jsonContent)
	}
	larkEventOperationLog := do.LarkEventOperationLog{
		Id:        uuid.New().String(),
		MsgInfo:   msgInfoJson,
		EventType: "msg",
	}
	_ = service.LarkEventOperationLog().Create(ctx, larkEventOperationLog)

	actions := []Action{
		&ProcessedUniqueAction{}, //避免重复处理
		&ProcessMentionAction{},  //判断机器人是否应该被调用
		&EmptyAction{},           //空消息处理
		&HelpAction{},            //帮助处理
		&QueryCustomerAction{},
		&SqlAgentAction{},
		&QueryRoomAction{},
		&CheckoutAction{},
		&MessageAction{}, //消息处理
	}
	chain(data, actions...)
	return nil
}

var _ MessageHandlerInterface = (*MessageHandler)(nil)

func NewMessageHandler(config config.LarkConfig) MessageHandlerInterface {
	return &MessageHandler{
		//sessionCache: services.GetSessionCache(),
		msgCache: GetMsgCache(),
		config:   config,
	}
}

func (m MessageHandler) judgeIfMentionMe(mention []*larkim.
	MentionEvent) bool {
	if len(mention) != 1 {
		return false
	}
	return *mention[0].Name == m.config.FeishuBotName
}
