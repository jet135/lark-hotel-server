package third_service

import (
	"context"
	"fmt"
	"lark-hotel-server/api/report"
	"lark-hotel-server/internal/consts"
	"lark-hotel-server/internal/lark/initialization"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/utils"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"github.com/samber/lo"
)

func SyncDocData(ctx context.Context, saveStayData bool, appToken string, tableId string, belongDate time.Time) {
	operationalData, err := QueryOperationalData(ctx, appToken, tableId)
	if err != nil || !operationalData.Success() || len(operationalData.Data.Items) == 0 {
		return
	}
	items := operationalData.Data.Items

	belongDate = utils.WithTimeAtStartOfDate(belongDate)
	dayBills, err := service.Bill().Query(ctx, do.Bill{BelongDate: belongDate})
	if err != nil {
		return
	}

	todayBillMap := lo.KeyBy(dayBills, func(item entity.Bill) string {
		return item.SourceId
	})

	bills := make([]do.Bill, 0, len(items))
	multiStays := make([]do.MultiStay, 0, len(items))

	earlyCheckoutBillIds := make([]string, 0, len(items))
	for _, item := range items {
		fields := item.Fields
		customerName := utils.TableTextField("客户名称", fields)
		if strings.TrimSpace(customerName) == "" {
			continue
		}

		doBill := new(do.Bill)
		bill, ok := todayBillMap[*item.RecordId]
		if ok {
			doBill.Id = bill.Id
			delete(todayBillMap, *item.RecordId)
		} else {
			doBill.Id = uuid.New().String()
		}

		numberOfNights := utils.TableFloatField("入住天数", fields)
		payTime := utils.TableTimeField("支付时间", fields)
		checkinTime := utils.TableTimeField("入住时间", fields)
		amount := utils.TableFloatField("合计", fields)
		phone := utils.TableTextField("电话号码", fields)
		expectedCheckoutTime := utils.TableTimeField("预期退房时间", fields)
		roomNumber := utils.TableTextField("房号", fields)

		if amount < 0 {
			earlyCheckoutBillIds = append(earlyCheckoutBillIds, doBill.Id.(string))
		}

		if numberOfNights == 0 {
			numberOfNights = 1
		}

		if saveStayData && numberOfNights > 1 && amount > 0 {
			multiStays = append(multiStays, do.MultiStay{
				CustomerName:         customerName,
				RoomNumber:           roomNumber,
				PayTime:              payTime,
				CheckinTime:          checkinTime,
				NumberOfNights:       numberOfNights,
				Done:                 false,
				BillId:               doBill.Id,
				ExpectedCheckoutTime: expectedCheckoutTime,
			})
		}

		doBill.Name = customerName
		doBill.RoomNumber = roomNumber
		doBill.Phone = phone
		doBill.Source = consts.SourceLark
		doBill.SourceId = *item.RecordId
		doBill.RoomPrice = utils.TableFloatField("房价", fields)
		doBill.RoomPaymentType = utils.TableStringField("房费方式", fields)
		doBill.Deposit = utils.TableFloatField("押金金额", fields)
		doBill.Amount = amount
		doBill.DepositType = utils.TableStringField("押金方式", fields)
		doBill.PayTime = payTime
		doBill.Shift = utils.TableStringField("班次", fields)
		doBill.Remark = utils.TableTextField("备注", fields)
		doBill.NumberOfNights = numberOfNights
		doBill.ExpectedCheckoutTime = expectedCheckoutTime
		doBill.CheckinTime = checkinTime
		doBill.CheckoutTime = utils.TableTimeField("退房时间", fields)
		doBill.DepositRefundTime = utils.TableTimeField("退押金时间", fields)
		doBill.BelongDate = belongDate
		doBill.IsAdditionalItem = utils.TableBoolField("附加收入/支出", fields)

		bills = append(bills, *doBill)
	}
	if len(todayBillMap) > 0 {
		todayBillIds := lo.Map(lo.Values(todayBillMap), func(todayBill entity.Bill, index int) string {
			return todayBill.Id
		})
		_ = service.Bill().Delete(ctx, todayBillIds)
	}

	if len(bills) == 0 {
		return
	}
	_ = service.Bill().BatchSave(ctx, &bills)

	// 记录multi_stay
	if saveStayData && len(multiStays) > 0 {
		billIds := lo.Map(multiStays, func(multiStay do.MultiStay, _ int) string {
			return multiStay.BillId.(string)
		})
		_ = service.MultiStay().DeleteByBillIds(ctx, billIds)
		err = service.MultiStay().Create(ctx, &multiStays)
		if err != nil {
			glog.Errorf(ctx, "create multiStay err: %v", err)
		}
	}

	if len(earlyCheckoutBillIds) > 0 {
		earlyCheckoutBills, err := service.Bill().ListByIds(ctx, &earlyCheckoutBillIds)
		if err == nil && len(earlyCheckoutBills) > 0 {
			for _, checkoutBill := range earlyCheckoutBills {
				_ = service.MultiStay().SettingDoneByRoomNumberAndCheckInTime(ctx, checkoutBill.RoomNumber, checkoutBill.CheckinTime)
			}

		}
	}

}

func StatisticsOperationalDataOfDate(ctx context.Context, date time.Time) (totalAmount float64, roomCount int) {
	date = utils.WithTimeAtStartOfDate(date)
	appInfo, err := service.LarkAppInfo().GetByBelongDate(ctx, date)
	if err != nil || appInfo == nil {
		return
	}
	operationalData, err := QueryOperationalData(ctx, appInfo.AppToken, appInfo.TableId)
	if err != nil || !operationalData.Success() || len(operationalData.Data.Items) == 0 {
		return
	}
	items := operationalData.Data.Items
	totalAmount = 0.0
	roomCount = 0

	withTimeAtEndOfDate := utils.WithTimeAtEndOfDate(date)
	for _, item := range items {
		fields := item.Fields
		amount := utils.TableFloatField("合计", fields)
		totalAmount += amount

		isAdditionalItem := utils.TableBoolField("附加收入/支出", fields)
		if isAdditionalItem {
			continue
		}

		checkinTime := utils.TableTimeField("入住时间", fields)
		// 判断入住时间是否大于date
		if checkinTime == nil || checkinTime.After(withTimeAtEndOfDate) {
			continue
		}

		roomCount++
	}
	return
}

func SendStatisticsOfDate(ctx context.Context, date time.Time, notify bool) {
	totalAmount, roomCount := StatisticsOperationalDataOfDate(ctx, date)

	larkConfig := initialization.GetLarkConfig()

	_, _ = createAppTableRecord(ctx, larkConfig.AnnualStatisticsAppToken, larkConfig.AnnualStatisticsTableId, map[string]interface{}{
		"日期":  fmt.Sprintf("%d/%d", date.Month(), date.Day()),
		"收入":  totalAmount,
		"开房数": roomCount,
		"月份":  fmt.Sprintf("%d月", date.Month()),
	})

	if notify {
		content := fmt.Sprintf("日期：%d月%d日 收入：%.2f元 开房数：%d间", date.Month(), date.Day(), totalAmount, roomCount)
		MessagesText(ctx, larkConfig.ManagerUserId, content)
	}
}

func SyncYesterdayBillData(ctx context.Context) {
	withTimeAtStartOfDay := utils.WithTimeAtStartOfDate(time.Now())
	// 获取前一天的app信息
	withTimeAtStartOfYesterday := withTimeAtStartOfDay.Add(-24 * time.Hour)
	larkAppInfo, err := service.LarkAppInfo().GetByBelongDate(ctx, withTimeAtStartOfYesterday)
	if err != nil || larkAppInfo == nil {
		return
	}

	SyncDocData(ctx, true, larkAppInfo.AppToken, larkAppInfo.TableId, withTimeAtStartOfYesterday)
}

func SendToLarkByBill(ctx context.Context, bill entity.Bill) (err error) {
	client := initialization.GetLarkClient()

	remark := bill.Remark
	if !strings.Contains(remark, "续住") {
		remark = strings.TrimSpace(fmt.Sprintf("续住 %s", remark))
	}

	fieldMap := map[string]interface{}{
		"客户名称": bill.Name,
		"房号":   bill.RoomNumber,
		"房价":   bill.RoomPrice,
		"合计":   0,
		"入住天数": bill.NumberOfNights,
		"房费方式": bill.RoomPaymentType,
		"押金方式": bill.DepositType,
		"押金金额": bill.Deposit,
		"电话号码": bill.Phone,
		"备注":   remark,
	}
	if !bill.PayTime.IsZero() {
		fieldMap["支付时间"] = bill.PayTime.Unix() * 1000
	}
	if !bill.CheckinTime.IsZero() {
		fieldMap["入住时间"] = bill.CheckinTime.Unix() * 1000
	}
	if !bill.ExpectedCheckoutTime.IsZero() {
		fieldMap["预期退房时间"] = bill.ExpectedCheckoutTime.Unix() * 1000
	}

	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(initialization.GetTodayAppToken()).
		TableId(initialization.GetTodayTableId()).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(fieldMap).
			Build()).
		Build()
	_, err = client.Bitable.AppTableRecord.Create(ctx, req)
	return
}

func SendToLark(reportData report.Data, ctx context.Context) (*larkbitable.CreateAppTableRecordResp, error) {

	fields := map[string]interface{}{
		"客户名称": reportData.Name,
		"房号":   reportData.RoomNumber,
		"入住时间": reportData.Date.Unix() * 1000,
		"电话号码": reportData.Phone,
		"班次":   reportData.Shift,
	}

	return createAppTableRecord(ctx, initialization.GetTodayAppToken(), initialization.GetTodayTableId(), fields)
}

func createAppTableRecord(ctx context.Context, appToken string, tableId string, fields map[string]interface{}) (*larkbitable.CreateAppTableRecordResp, error) {
	client := initialization.GetLarkClient()
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableId).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(fields).
			Build()).
		Build()
	resp, err := client.Bitable.AppTableRecord.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateLarkCheckoutInfo(ctx context.Context, bill entity.Bill, appToken string, tableId string, checkoutTime time.Time) (*larkbitable.UpdateAppTableRecordResp, error) {
	client := initialization.GetLarkClient()

	checkoutUnix := checkoutTime.Unix() * 1000
	fields := map[string]interface{}{
		"退房时间": checkoutUnix,
	}
	if bill.Deposit > 0 {
		fields["退押金时间"] = checkoutUnix
	}

	req := larkbitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableId).
		RecordId(bill.SourceId).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(fields).
			Build()).
		Build()
	return client.Bitable.AppTableRecord.Update(ctx, req)
}

type AppTableResp struct {
	AppToken    string `json:"app_token,omitempty"` // 多维表格 app token
	TableId     string
	Name        string `json:"name,omitempty"`         // 多维表格 App 名字
	FolderToken string `json:"folder_token,omitempty"` // 多维表格 App 归属文件夹
	Url         string `json:"url,omitempty"`          // 多维表格 App URL
}

func GetOrElseGenerateOperationalDataApp(ctx context.Context, belongDate time.Time) (*AppTableResp, error) {

	todayApp, _ := service.LarkAppInfo().GetByBelongDate(ctx, belongDate)

	larkConfig := initialization.GetLarkConfig()

	if todayApp != nil {
		return &AppTableResp{
			AppToken:    todayApp.AppToken,
			TableId:     todayApp.TableId,
			Name:        todayApp.AppName,
			FolderToken: todayApp.FolderToken,
		}, nil
	}

	appName := fmt.Sprintf("%s月%s日-运营数据", belongDate.Format("1"), belongDate.Format("2"))
	app, err := CopyApp(appName, larkConfig.WordTemplateAppToken, larkConfig.WordAppFolderToken)
	if err != nil || app == nil {
		return nil, err
	}
	todayApp = &entity.LarkAppInfo{
		AppName:     app.Name,
		AppToken:    app.AppToken,
		FolderToken: app.FolderToken,
		Url:         app.Url,
		BelongDate:  belongDate,
		TableId:     app.TableId,
	}
	_ = service.LarkAppInfo().Create(ctx, todayApp)

	return app, nil
}

func CopyApp(name string, appToken string, folderToken string) (appTableResp *AppTableResp, err error) {
	client := initialization.GetLarkClient()

	req := larkbitable.NewCopyAppReqBuilder().
		AppToken(appToken).
		Body(larkbitable.NewCopyAppReqBodyBuilder().
			Name(name).
			FolderToken(folderToken).
			WithoutContent(false).
			TimeZone("Asia/Shanghai").
			Build()).
		Build()
	ctx := context.Background()
	copyAppResp, err := client.Bitable.App.Copy(ctx, req)
	if err != nil || !copyAppResp.Success() {
		return nil, err
	}

	// 等三秒创建完
	time.Sleep(3 * time.Second)

	// 获取运营的tableId
	appInfo := copyAppResp.Data.App
	appTableResp = new(AppTableResp)
	appToken = *appInfo.AppToken
	appTableResp.AppToken = appToken
	appTableResp.Name = *appInfo.Name
	appTableResp.FolderToken = *appInfo.FolderToken
	appTableResp.Url = *appInfo.Url

	tableResp, err := ListAppTable(appToken)
	if err != nil || !tableResp.Success() {
		return nil, err
	}

	for _, item := range tableResp.Data.Items {
		if *item.Name == initialization.GetLarkConfig().WordTemplateTableName {
			appTableResp.TableId = *item.TableId
			break
		}
	}

	// 删掉复制后的数据
	searchResp, err := QueryOperationalData(ctx, appToken, appTableResp.TableId)

	// 处理错误
	if err != nil || !searchResp.Success() {
		return
	}

	recordIds := make([]string, 0, 500)
	for _, item := range searchResp.Data.Items {
		recordIds = append(recordIds, *item.RecordId)
	}

	deleteReq := larkbitable.NewBatchDeleteAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(appTableResp.TableId).
		Body(larkbitable.NewBatchDeleteAppTableRecordReqBodyBuilder().
			Records(recordIds).
			Build()).
		Build()

	_, _ = client.Bitable.AppTableRecord.BatchDelete(ctx, deleteReq)

	return
}

func ListAppTable(appToken string) (*larkbitable.ListAppTableResp, error) {
	client := initialization.GetLarkClient()
	req := larkbitable.NewListAppTableReqBuilder().
		AppToken(appToken).
		Build()
	return client.Bitable.AppTable.List(context.Background(), req)
}

func QueryOperationalData(ctx context.Context, appToken string, tableId string) (*larkbitable.SearchAppTableRecordResp, error) {
	client := initialization.GetLarkClient()
	req := larkbitable.NewSearchAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableId).
		PageSize(500).
		Body(larkbitable.NewSearchAppTableRecordReqBodyBuilder().
			Build()).
		Build()
	return client.Bitable.AppTableRecord.Search(ctx, req)
}
