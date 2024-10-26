package lark

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

func doProcess(resp *ghttp.Response, req *ghttp.Request, reqHandler larkevent.IReqHandler, options ...larkevent.OptionFunc) {
	// 转换http请求对象为标准请求对象
	ctx := context.Background()
	eventReq, err := translate(ctx, req)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}

	//处理请求
	eventResp := reqHandler.Handle(ctx, eventReq)

	// 回写结果
	err = write(ctx, resp, eventResp)
	if err != nil {
		reqHandler.Logger().Error(ctx, fmt.Sprintf("write resp result error:%s", err.Error()))
	}
}

func NewCardActionHandlerFunc(cardActionHandler *larkcard.CardActionHandler, options ...larkevent.OptionFunc) func(r *ghttp.Request) {

	// 构建模板类
	cardActionHandler.InitConfig(options...)
	return func(r *ghttp.Request) {
		doProcess(r.Response, r, cardActionHandler, options...)
	}
}

func NewEventHandlerFunc(eventDispatcher *dispatcher.EventDispatcher, options ...larkevent.OptionFunc) func(r *ghttp.Request) {
	eventDispatcher.InitConfig(options...)
	return func(r *ghttp.Request) {
		doProcess(r.Response, r, eventDispatcher, options...)
	}
}

func processError(ctx context.Context, logger larkcore.Logger, path string, err error) *larkevent.EventResp {
	header := map[string][]string{}
	header[larkevent.ContentTypeHeader] = []string{larkevent.DefaultContentType}
	eventResp := &larkevent.EventResp{
		Header:     header,
		Body:       []byte(fmt.Sprintf(larkevent.WebhookResponseFormat, err.Error())),
		StatusCode: http.StatusInternalServerError,
	}
	logger.Error(ctx, fmt.Sprintf("event handle err:%s, %v", path, err))
	return eventResp
}

func write(ctx context.Context, resp *ghttp.Response, eventResp *larkevent.EventResp) error {
	resp.WriteHeader(eventResp.StatusCode)
	for k, vs := range eventResp.Header {
		for _, v := range vs {
			resp.Header().Add(k, v)
		}
	}

	if len(eventResp.Body) > 0 {
		_, err := resp.BufferWriter.Write(eventResp.Body)
		return err
	}
	return nil
}
func translate(ctx context.Context, req *ghttp.Request) (*larkevent.EventReq, error) {
	rawBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	eventReq := &larkevent.EventReq{
		Header:     req.Header,
		Body:       rawBody,
		RequestURI: req.RequestURI,
	}

	return eventReq, nil
}
