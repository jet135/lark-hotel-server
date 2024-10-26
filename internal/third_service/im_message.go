package third_service

import (
	"context"
	"errors"
	"lark-hotel-server/internal/lark/initialization"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func MessagesText(ctx context.Context,
	receiveId string,
	textContent string,
) error {
	if receiveId == "" {
		return nil
	}
	client := initialization.GetLarkClient()
	content := larkim.NewTextMsgBuilder().
		Text(textContent).
		Build()
	resp, err := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("user_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType(larkim.MsgTypeText).
			Uuid(uuid.New().String()).
			Content(content).
			Build()).
		Build())

	// 处理错误
	if err != nil {
		glog.Error(ctx, err)
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		glog.Errorf(ctx, "resp.Code:%v, resp.Msg:%s, resp.RequestId:%s", resp.Code, resp.Msg, resp.RequestId())
		return errors.New(resp.Msg)
	}
	return nil
}
