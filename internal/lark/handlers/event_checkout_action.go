package handlers

import (
	"context"
	"fmt"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/third_service"
	"lark-hotel-server/internal/utils"
	"strings"
	"time"

	"github.com/samber/lo"
)

type CheckoutAction struct { /*退房*/
}

func (*CheckoutAction) Execute(a *ActionInfo) bool {
	if checkoutKey, found := utils.EitherCutPrefix(a.info.qParsed,
		"tf", "退房"); found {
		ctx := *a.ctx
		msgId := *a.info.msgId

		checkoutKey = strings.TrimSpace(checkoutKey)
		split := strings.Split(checkoutKey, " ")
		if len(split) > 1 {
			replyMsg(ctx, "🤖️：输入格式有误", msgId)
		}
		roomNumber := split[0]

		now := time.Now()
		withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(now)
		withTimeAtStartOfYesterday := withTimeAtStartOfDay.Add(-24 * time.Hour)
		bills, err := service.Bill().Query(ctx, do.Bill{RoomNumber: roomNumber, BelongDate: withTimeAtStartOfYesterday})
		if err != nil {
			replyMsg(ctx, "🤖️：退房失败，请联系管理员", msgId)
			return false
		}
		yesterDayNoneRoomNumber := false
		if len(bills) > 0 {
			yesterDayNoneRoomNumber = true
			bills = lo.Filter(bills, func(item entity.Bill, _ int) bool {
				return item.CheckoutTime.IsZero()
			})
		}

		if len(bills) > 0 {
			larkAppInfo, err := service.LarkAppInfo().GetByBelongDate(ctx, withTimeAtStartOfYesterday)
			if err != nil || larkAppInfo == nil {
				checkoutCommonErrSend(ctx, msgId)
				return false
			}
			for _, bill := range bills {
				_, err := third_service.UpdateLarkCheckoutInfo(ctx, bill, larkAppInfo.AppToken, larkAppInfo.TableId, now)
				if err != nil {
					checkoutCommonErrSend(ctx, msgId)
					return false
				}
			}
		} else {
			bills, err = service.Bill().Query(ctx, do.Bill{RoomNumber: roomNumber, BelongDate: withTimeAtStartOfDay})
			if err != nil {
				checkoutCommonErrSend(ctx, msgId)
				return false
			}
			if len(bills) > 0 {
				bills = lo.Filter(bills, func(item entity.Bill, _ int) bool {
					return item.CheckoutTime.IsZero()
				})
			} else {
				if yesterDayNoneRoomNumber {
					replyMsg(ctx, "🤖️：查不到该房间信息，请核查房间号", msgId)
					return false
				}
			}
			if len(bills) > 0 {
				larkAppInfo, err := service.LarkAppInfo().GetByBelongDate(ctx, withTimeAtStartOfDay)
				if err != nil {
					checkoutCommonErrSend(ctx, msgId)
					return false
				}
				for _, bill := range bills {
					_, err := third_service.UpdateLarkCheckoutInfo(ctx, bill, larkAppInfo.AppToken, larkAppInfo.TableId, now)
					if err != nil {
						checkoutCommonErrSend(ctx, msgId)
						return false
					}
				}
			} else {
				replyMsg(ctx, "🤖️：不可重复退房", msgId)
				return false
			}
		}

		// 飞书消息通知和更改bill信息
		billOfUpdates := make([]do.Bill, 0, len(bills))
		for _, bill := range bills {
			billOfUpdate := do.Bill{
				Id:           bill.Id,
				CheckoutTime: now,
			}
			depositStr := lo.TernaryF(bill.Deposit > 0, func() string {
				billOfUpdate.DepositRefundTime = now
				return fmt.Sprintf("%.2f %s", bill.Deposit, bill.DepositType)
			}, func() string {
				return fmt.Sprintf("%.2f", bill.Deposit)
			})
			billOfUpdates = append(billOfUpdates, billOfUpdate)
			replyMsg(ctx, fmt.Sprintf("🤖️：%s已退房，退房信息 [客户名称：%s｜押金：%s｜入住时间：%s｜退房时间：%s]", bill.RoomNumber, bill.Name, depositStr, utils.FormatTime(bill.CheckinTime), utils.FormatTime(now)), msgId)
		}
		_ = service.Bill().BatchSave(ctx, &billOfUpdates)
		return false
	}
	return true
}

func checkoutCommonErrSend(ctx context.Context, msgId string) {
	replyMsg(ctx, "🤖️：退房失败，请联系管理员", msgId)
}
