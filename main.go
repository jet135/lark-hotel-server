package main

import (
	"context"
	"lark-hotel-server/internal/cmd"
	"lark-hotel-server/internal/config"
	"lark-hotel-server/internal/consts"
	"lark-hotel-server/internal/lark/handlers"
	"lark-hotel-server/internal/lark/initialization"
	_ "lark-hotel-server/internal/logic"
	_ "lark-hotel-server/internal/packed"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/third_service"
	"lark-hotel-server/internal/utils"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

var isTimingSyncDocData bool

func main() {
	initLark()
	scheduled()
	cmd.Main.Run(gctx.GetInitCtx())
}

func scheduled() {
	_, err := gcron.Add(gctx.New(), "# 30 23 * * *", func(ctx context.Context) {
		withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(time.Now())
		tomorrow := withTimeAtStartOfDay.Add(24 * time.Hour)

		if _, err := third_service.GetOrElseGenerateOperationalDataApp(ctx, tomorrow); err != nil {
			glog.Errorf(ctx, "generateAndSettingLarkBaseData err: %v", err)
		}
	})
	if err != nil {
		panic(err)
	}

	if isTimingSyncDocData {
		// 9min更新一次昨日数据
		_, err = gcron.Add(gctx.New(), "# */9 * * * *", func(ctx context.Context) {
			withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(time.Now())
			withTimeAtStartOfYesterday := withTimeAtStartOfDay.Add(-24 * time.Hour)
			if larkAppInfo, err2 := service.LarkAppInfo().GetByBelongDate(ctx, withTimeAtStartOfYesterday); err2 == nil {
				third_service.SyncDocData(ctx, false, larkAppInfo.AppToken, larkAppInfo.TableId, withTimeAtStartOfYesterday)
			}
		})
		if err != nil {
			panic(err)
		}
		// 5min更新一次数据
		_, err = gcron.Add(gctx.New(), "# */5 * * * *", func(ctx context.Context) {
			third_service.SyncDocData(ctx, false, initialization.GetTodayAppToken(), initialization.GetTodayTableId(), time.Now())
		})
		if err != nil {
			panic(err)
		}
		_, err = gcron.Add(gctx.New(), "# 0 8 * * *", func(ctx context.Context) {
			syncYesterdayBillData(ctx)

			// 同步昨日数据统计
			withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(time.Now())
			withTimeAtStartOfYesterday := withTimeAtStartOfDay.Add(-24 * time.Hour)
			third_service.SendStatisticsOfDate(ctx, withTimeAtStartOfYesterday, true)
		})
		if err != nil {
			panic(err)
		}
	}

	_, err = gcron.Add(gctx.New(), "# 0 0 * * *", func(ctx context.Context) {
		_ = generateAndSettingLarkBaseData(ctx)
	})
	if err != nil {
		panic(err)
	}
}

func syncYesterdayBillData(ctx context.Context) {
	third_service.SyncYesterdayBillData(ctx)

	multiStays, err := service.MultiStay().FindByDone(ctx, consts.BusinessFlagNo)
	if err != nil {
		return
	}
	if multiStays == nil || len(multiStays) == 0 {
		return
	}
	withTimeAtEndOfToday := utils.WithTimeAtEndOfDate(time.Now())

	billIds := make([]string, 0, len(multiStays))
	for _, multiStay := range multiStays {
		if multiStay.CheckinTime.After(withTimeAtEndOfToday) {
			continue
		}
		// 判断是否生成第二天的账单
		needGenerate := false
		if !multiStay.ExpectedCheckoutTime.IsZero() {
			if multiStay.ExpectedCheckoutTime.After(withTimeAtEndOfToday) {
				needGenerate = true
			}
		} else {
			var remainingTime = int64(24 * multiStay.NumberOfNights)
			stayTime := multiStay.CheckinTime.Add(time.Duration(remainingTime * int64(time.Hour)))
			if stayTime.After(withTimeAtEndOfToday) {
				needGenerate = true
			}
		}
		if needGenerate {
			billIds = append(billIds, multiStay.BillId)
		} else {
			_ = service.MultiStay().SettingDone(ctx, multiStay.Id)
		}
	}
	bills, _ := service.Bill().ListByIds(ctx, &billIds)
	if bills != nil && len(bills) > 0 {
		for _, bill := range bills {
			// todo 如果有押金，设置前天bill的退押金时间（等于把昨天的押金移到今天）
			_ = third_service.SendToLarkByBill(ctx, bill)
		}
	}
}

func initLark() {
	var larkConfig config.LarkConfig
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "lark").Struct(&larkConfig)
	if err != nil {
		panic("init config error")
	}
	isTimingSyncDocData = larkConfig.TimingSyncDocData
	initialization.LoadLarkClient(larkConfig)
	handlers.InitHandlers(larkConfig)

	if err := generateAndSettingLarkBaseData(ctx); err != nil {
		panic("setting app base config err")
	}
}

func generateAndSettingLarkBaseData(ctx context.Context) error {
	withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(time.Now())
	todayApp, err := third_service.GetOrElseGenerateOperationalDataApp(ctx, withTimeAtStartOfDay)
	if err != nil || todayApp == nil {
		return err
	}

	initialization.LoadTodayAppToken(todayApp.AppToken)
	initialization.LoadTodayTableId(todayApp.TableId)

	return nil
}
