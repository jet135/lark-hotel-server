package handlers

import (
	"lark-hotel-server/internal/model/bo"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/utils"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
)

type QueryRoomAction struct { /*查询房间*/
}

func (*QueryRoomAction) Execute(a *ActionInfo) bool {
	if roomKey, found := utils.EitherCutPrefix(a.info.qParsed,
		"fh", "房号"); found {
		roomKey = strings.TrimSpace(roomKey)
		ctx := *a.ctx
		billQue := do.Bill{RoomNumber: roomKey}

		bills, err := service.Bill().Query(ctx, billQue)
		if err == nil && bills != nil {
			allBills := lo.Filter(bills, func(item entity.Bill, _ int) bool {
				return item.Amount >= 0
			})
			paymentBills := lo.Filter(allBills, func(item entity.Bill, _ int) bool {
				return item.Amount > 0
			})

			sort.Slice(paymentBills, func(i, j int) bool {
				return paymentBills[i].PayTime.After(paymentBills[j].PayTime)
			})

			boBills := make([]bo.Bill, 0, len(paymentBills))

			boRoomBill := bo.RoomBill{}
			firstBill := paymentBills[0]
			boRoomBill.RoomNumber = roomKey
			boRoomBill.LastCheckInTime = firstBill.CheckinTime
			boRoomBill.TotalCheckIns = len(allBills)

			withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(time.Now())
			time30DaysAgo := withTimeAtStartOfDay.Add(-29 * 24 * time.Hour)
			boRoomBill.CheckInsLast30Days = len(lo.Filter(allBills, func(item entity.Bill, _ int) bool {
				return item.CheckinTime.After(time30DaysAgo)
			}))

			var HighestRoomPrice float64
			var LowestRoomPrice = math.MaxFloat64
			lo.ForEach(paymentBills, func(item entity.Bill, _ int) {
				if item.RoomPrice > HighestRoomPrice {
					HighestRoomPrice = item.RoomPrice
				}
				if item.RoomPrice < LowestRoomPrice {
					LowestRoomPrice = item.RoomPrice
				}
			})
			boRoomBill.LowestRoomPrice = LowestRoomPrice
			boRoomBill.HighestRoomPrice = HighestRoomPrice

			for _, customerBill := range paymentBills {
				boBill := new(bo.Bill)
				err := gconv.Struct(customerBill, boBill)
				if err != nil {
					customerCommonReplyErrMsg(ctx, *a.info.msgId)
					return false
				}
				boBills = append(boBills, *boBill)
			}
			boRoomBill.Bills = boBills
			boRoomBill.BillTotal = len(boBills)
			sendRoomInfoCard(ctx, a.info.msgId, boRoomBill)
		} else {
			replyMsg(ctx, "🤖️：未找到该房间数据", *a.info.msgId)
		}
		return false
	}
	return true
}
