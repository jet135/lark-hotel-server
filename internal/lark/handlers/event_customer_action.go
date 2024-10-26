package handlers

import (
	"context"
	"lark-hotel-server/internal/model/bo"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/utils"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
)

type QueryCustomerAction struct { /*查询客户*/
}

func (*QueryCustomerAction) Execute(a *ActionInfo) bool {
	if customerKey, found := utils.EitherCutPrefix(a.info.qParsed,
		"kh", "客户"); found {
		customerKey = strings.TrimSpace(customerKey)
		strDefine := utils.StrDefine(customerKey)
		ctx := *a.ctx
		var billQue do.Bill
		if strDefine == utils.StrIsName {
			billQue = do.Bill{Name: customerKey}
		} else if strDefine == utils.StrIsPhone {
			billQue = do.Bill{Phone: customerKey}
		} else {
			replyMsg(ctx, "🤖️：输入格式有误", *a.info.msgId)
			return false
		}

		bills, err := service.Bill().QueryPayment(ctx, billQue)
		if err == nil && bills != nil {
			sort.Slice(bills, func(i, j int) bool {
				return bills[i].PayTime.After(bills[j].PayTime)
			})
			billGroup := lo.GroupBy(bills, func(bill entity.Bill) string {
				return utils.UniqueBillKey(bill.Name, bill.Phone)
			})
			for _, customerBills := range billGroup {
				boBills := make([]bo.Bill, 0, len(customerBills))

				boCustomer := bo.Customer{}
				firstBill := customerBills[0]
				boCustomer.Name = firstBill.Name
				boCustomer.Phone = firstBill.Phone
				boCustomer.LastCheckInTime = firstBill.CheckinTime

				for _, customerBill := range customerBills {
					boBill := new(bo.Bill)
					err := gconv.Struct(customerBill, boBill)
					if err != nil {
						customerCommonReplyErrMsg(ctx, *a.info.msgId)
						return false
					}
					boBills = append(boBills, *boBill)
				}
				boCustomer.Bills = boBills
				boCustomer.BillTotal = len(boBills)
				sendCustomerInfoCard(ctx, a.info.msgId, boCustomer)
			}
		} else {
			replyMsg(ctx, "🤖️：未找到该客户", *a.info.msgId)
		}
		return false
	}
	return true
}

func customerCommonReplyErrMsg(ctx context.Context, msgId string) {
	replyMsg(ctx, "🤖️：查询客户信息故障了，我不会修，请管理员看看", msgId)
}
