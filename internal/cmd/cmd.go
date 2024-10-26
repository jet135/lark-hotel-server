package cmd

import (
	"context"
	"fmt"
	"lark-hotel-server/api/bill"
	"lark-hotel-server/api/report"
	"lark-hotel-server/internal/config"
	"lark-hotel-server/internal/lark"
	"lark-hotel-server/internal/lark/handlers"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/third_service"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

var apiToken string

func init() {
	var serverConfig config.ServerConfig
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "server").Struct(&serverConfig)
	if err != nil {
		panic("init server config error")
	}
	apiToken = serverConfig.ApiToken
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server()
			s.Use(ghttp.MiddlewareCORS)
			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.Write("hello world")
			})
			s.BindHandler("/gdzwfw/report", func(r *ghttp.Request) {
				rCtx := r.Context()
				var req report.Req
				if err := r.Parse(&req); err != nil {
					g.Log().Error(rCtx, "report param error..")
					r.Exit()
				}
				if err := service.ReportLogs().Process(rCtx, req); err != nil {
					g.Log().Error(rCtx, "report process error..")
				}
				r.Response.WriteJson(ghttp.DefaultHandlerResponse{
					Code:    http.StatusOK,
					Message: "ok",
				})
			})

			s.BindHandler("/api/statistics", func(r *ghttp.Request) {
				authorization := r.GetHeader("Authorization")
				token := strings.TrimPrefix(authorization, "Bearer ")

				if token == "" || apiToken != token {
					r.Response.WriteJson(ghttp.DefaultHandlerResponse{
						Code: http.StatusUnauthorized,
						Data: "api_token有误",
					})
					return
				}
				rCtx := r.Context()
				var req bill.StatisticsReq
				if err := r.Parse(&req); err != nil {
					g.Log().Error(rCtx, "report param error..")
					r.Exit()
				}

				layout := "2006年1月2日 15:04"

				// 解析时间字符串
				t, err := time.ParseInLocation(layout, req.Date, time.Now().Location())
				if err != nil {
					r.Response.WriteJson(ghttp.DefaultHandlerResponse{
						Code:    http.StatusBadRequest,
						Message: "解析时间出错",
						Data:    "解析时间出错",
					})
					return
				}

				totalAmount, roomCount := third_service.StatisticsOperationalDataOfDate(rCtx, t)
				msg := fmt.Sprintf("日期：%d月%d日 收入：%.2f元 开房数：%d间", t.Month(), t.Day(), totalAmount, roomCount)
				r.Response.WriteJson(ghttp.DefaultHandlerResponse{
					Code:    http.StatusOK,
					Message: "ok",
					Data:    msg,
				})
			})

			eventHandler := dispatcher.NewEventDispatcher("", "").
				OnP2MessageReceiveV1(handlers.Handler)
			s.BindHandler("/webhook/event", lark.NewEventHandlerFunc(eventHandler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))
			s.Run()
			return nil
		},
	}
)
