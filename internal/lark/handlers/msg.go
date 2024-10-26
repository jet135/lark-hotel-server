package handlers

import (
	"context"
	"errors"
	"fmt"
	"lark-hotel-server/internal/consts"
	"lark-hotel-server/internal/lark/initialization"
	"lark-hotel-server/internal/model/bo"
	"lark-hotel-server/internal/utils"
	"math"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/google/uuid"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/samber/lo"
)

func newSendCard(
	header *larkcard.MessageCardHeader,
	elements ...larkcard.MessageCardElement) string {
	config := larkcard.NewMessageCardConfig().
		WideScreenMode(false).
		EnableForward(true).
		UpdateMulti(false).
		Build()
	var aElementPool []larkcard.MessageCardElement
	for _, element := range elements {
		aElementPool = append(aElementPool, element)
	}
	// 卡片消息体
	cardContent, _ := larkcard.NewMessageCard().
		Config(config).
		Header(header).
		Elements(
			aElementPool,
		).
		String()
	return cardContent
}

func compareRoomTypePrice(roomNumber string, roomPrice float64, minRoomPrices map[consts.RoomType]float64, maxRoomPrices map[consts.RoomType]float64) {
	roomType := utils.JudgeRoomType(roomNumber)
	if _, exists := minRoomPrices[roomType]; !exists {
		minRoomPrices[roomType] = math.MaxInt
	}
	if roomPrice < minRoomPrices[roomType] {
		minRoomPrices[roomType] = roomPrice
	}
	if roomPrice > maxRoomPrices[roomType] {
		maxRoomPrices[roomType] = roomPrice
	}
}

func CustomerInfoCardTemplate(customer bo.Customer) string {
	//billList := []map[string]string{
	//	{"pay_time": "🎯 支付时间：2006-01-02 15:04", "bill_info": "房号：803｜房价：188 微信｜押金：100 现金｜合计：188｜入住天数：1｜入住时间：2024/07/08 12:00｜退房时间：2024/07/09 14:00｜退押金时间：2024/07/09 14:00"},
	//	{"pay_time": "🎯 支付时间：2006-01-02 15:04", "bill_info": "房号：803｜房价：188 微信｜押金：100 现金｜合计：188｜入住天数：1｜入住时间：2024/07/08 12:00｜退房时间：2024/07/09 14:00｜退押金时间：2024/07/09 14:00"},
	//	{"pay_time": "🎯 支付时间：2006-01-02 15:04", "bill_info": "房号：803｜房价：188 微信｜押金：100 现金｜合计：188｜入住天数：1｜入住时间：2024/07/08 12:00｜退房时间：2024/07/09 14:00｜退押金时间：2024/07/09 14:00"},
	//}
	//templateVariable := map[string]interface{}{
	//	"customerName":     "张三",
	//	"customerInfo":     "姓名：张三｜手机号：135555555｜上次入住时间：2024/07/08 12:00",
	//	"lowestRoomPrice":  "单床房：128｜双床房：148｜小套房：188｜小麻将套房：218｜大套房：288",
	//	"highestRoomPrice": "单床房：168｜双床房：188｜小套房：218｜小麻将套房：288｜大套房：388",
	//	"billTotal":        "3",
	//	"billList":         billList,
	//}
	minRoomPrices := make(map[consts.RoomType]float64, 6)
	maxRoomPrices := make(map[consts.RoomType]float64, 6)

	templateVariable := make(map[string]interface{}, 6)
	templateVariable["customerName"] = customer.Name
	templateVariable["customerInfo"] = fmt.Sprintf("手机号：%s｜上次入住时间：%s", customer.Phone, utils.FormatTime(customer.LastCheckInTime))
	templateVariable["billTotal"] = customer.BillTotal

	if customer.Bills != nil && len(customer.Bills) > 0 {
		billList := make([]map[string]string, 0, len(customer.Bills))
		for _, bill := range customer.Bills {
			compareRoomTypePrice(bill.RoomNumber, bill.RoomPrice, minRoomPrices, maxRoomPrices)

			numberOfNights := 1
			if bill.NumberOfNights > 0 {
				numberOfNights = bill.NumberOfNights
			}

			checkoutTimeStr := lo.Ternary(bill.CheckoutTime.IsZero(), "-", utils.FormatTime(bill.CheckoutTime))
			depositRefundTimeStr := lo.Ternary(bill.DepositRefundTime.IsZero(), "-", utils.FormatTime(bill.DepositRefundTime))

			billList = append(billList, map[string]string{
				"pay_time": fmt.Sprintf("🎯 支付时间：%s", utils.FormatTime(bill.PayTime)),
				"bill_info": fmt.Sprintf("房号：%s｜房价：%s｜押金：%s｜合计：%.2f｜入住天数：%d｜入住时间：%s｜退房时间：%s｜退押金时间：%s",
					bill.RoomNumber, fmt.Sprintf("%.2f %s", bill.RoomPrice, bill.RoomPaymentType), fmt.Sprintf("%.2f %s", bill.Deposit, bill.DepositType), bill.Amount, numberOfNights, utils.FormatTime(bill.CheckinTime), checkoutTimeStr, depositRefundTimeStr),
			})
		}

		templateVariable["billList"] = billList
	}

	minRoomPriceStrs := map[consts.RoomType]string{
		consts.SingleBed:         "-",
		consts.DoubleBed:         "-",
		consts.JuniorSuite:       "-",
		consts.SmallMahjongSuite: "-",
		consts.LargeSuite:        "-",
	}
	maxRoomPriceStrs := map[consts.RoomType]string{
		consts.SingleBed:         "-",
		consts.DoubleBed:         "-",
		consts.JuniorSuite:       "-",
		consts.SmallMahjongSuite: "-",
		consts.LargeSuite:        "-",
	}
	if len(minRoomPrices) > 0 {
		for roomType, price := range minRoomPrices {
			minRoomPriceStrs[roomType] = fmt.Sprintf("%.2f", price)
		}
	}
	if len(maxRoomPrices) > 0 {
		for roomType, price := range maxRoomPrices {
			maxRoomPriceStrs[roomType] = fmt.Sprintf("%.2f", price)
		}
	}
	templateVariable["lowestRoomPrice"] = fmt.Sprintf("单床房：%s｜双床房：%s｜小套房：%s｜小麻将套房：%s｜大套房：%s", minRoomPriceStrs[consts.SingleBed], minRoomPriceStrs[consts.DoubleBed], minRoomPriceStrs[consts.JuniorSuite], minRoomPriceStrs[consts.SmallMahjongSuite], minRoomPriceStrs[consts.LargeSuite])
	templateVariable["highestRoomPrice"] = fmt.Sprintf("单床房：%s｜双床房：%s｜小套房：%s｜小麻将套房：%s｜大套房：%s", maxRoomPriceStrs[consts.SingleBed], maxRoomPriceStrs[consts.DoubleBed], maxRoomPriceStrs[consts.JuniorSuite], maxRoomPriceStrs[consts.SmallMahjongSuite], maxRoomPriceStrs[consts.LargeSuite])

	return utils.NewCardTemplateCreateReqParam("AAqHGZ41L71bu", templateVariable)
}

func RoomInfoCardTemplate(roomBill bo.RoomBill) string {
	//billList := []map[string]string{
	//	{"pay_time": "🎯 支付时间：2006-01-02 15:04", "bill_info": "房号：803｜房价：188 微信｜押金：100 现金｜合计：188｜入住天数：1｜入住时间：2024/07/08 12:00｜退房时间：2024/07/09 14:00｜退押金时间：2024/07/09 14:00"},
	//	{"pay_time": "🎯 支付时间：2006-01-02 15:04", "bill_info": "房号：803｜房价：188 微信｜押金：100 现金｜合计：188｜入住天数：1｜入住时间：2024/07/08 12:00｜退房时间：2024/07/09 14:00｜退押金时间：2024/07/09 14:00"},
	//	{"pay_time": "🎯 支付时间：2006-01-02 15:04", "bill_info": "房号：803｜房价：188 微信｜押金：100 现金｜合计：188｜入住天数：1｜入住时间：2024/07/08 12:00｜退房时间：2024/07/09 14:00｜退押金时间：2024/07/09 14:00"},
	//}
	//templateVariable := map[string]interface{}{
	//	"customerName":     "张三",
	//	"roomInfo":     "总开房数：100｜近30天开房数：28｜上次入住时间：2024/07/08 12:00",
	//	"lowestRoomPrice":  "128",
	//	"highestRoomPrice": "168",
	//	"billTotal":        "3",
	//	"billList":         billList,
	//}

	templateVariable := make(map[string]interface{}, 6)
	templateVariable["roomNumber"] = roomBill.RoomNumber
	templateVariable["roomInfo"] = fmt.Sprintf("总开房数：%d｜近30天开房数：%d｜上次入住时间：%s", roomBill.TotalCheckIns, roomBill.CheckInsLast30Days, utils.FormatTime(roomBill.LastCheckInTime))
	templateVariable["billTotal"] = roomBill.BillTotal
	templateVariable["lowestRoomPrice"] = roomBill.LowestRoomPrice
	templateVariable["highestRoomPrice"] = roomBill.HighestRoomPrice

	if roomBill.Bills != nil && len(roomBill.Bills) > 0 {
		billList := make([]map[string]string, 0, len(roomBill.Bills))
		for _, bill := range roomBill.Bills {

			numberOfNights := 1
			if bill.NumberOfNights > 0 {
				numberOfNights = bill.NumberOfNights
			}

			checkoutTimeStr := lo.Ternary(bill.CheckoutTime.IsZero(), "-", utils.FormatTime(bill.CheckoutTime))
			depositRefundTimeStr := lo.Ternary(bill.DepositRefundTime.IsZero(), "-", utils.FormatTime(bill.DepositRefundTime))

			billList = append(billList, map[string]string{
				"pay_time": fmt.Sprintf("🎯 支付时间：%s", utils.FormatTime(bill.PayTime)),
				"bill_info": fmt.Sprintf("客户名称：%s｜房价：%s｜押金：%s｜合计：%.2f｜入住天数：%d｜入住时间：%s｜退房时间：%s｜退押金时间：%s",
					bill.Name, fmt.Sprintf("%.2f %s", bill.RoomPrice, bill.RoomPaymentType), fmt.Sprintf("%.2f %s", bill.Deposit, bill.DepositType), bill.Amount, numberOfNights, utils.FormatTime(bill.CheckinTime), checkoutTimeStr, depositRefundTimeStr),
			})
		}

		templateVariable["billList"] = billList
	}

	return utils.NewCardTemplateCreateReqParam("AAqHZaVbn5rLl", templateVariable)
}

func HelpCardTemplate() string {
	return newSendCard(
		withHeader("🎒我是小助手", larkcard.TemplateBlue),
		withMainMd("🤖 **查询客户信息** \n 文本回复 *客户* 或 *kh* 后面接客户姓名或手机号"),
		withSplitLine(),
		withMainMd("🛖 **查询房间信息** \n 文本回复 *房号* 或 *fh* 后面接房号"),
		withSplitLine(),
		withMainMd("🥷 **快捷退房**\n 文本回复 *退房* 或 *tf* 后面接房号"),
	)
}

func sendHelpCard(ctx context.Context, msgId *string) {
	replyCard(ctx, msgId, HelpCardTemplate())
}

func sendCustomerInfoCard(ctx context.Context, msgId *string, customer bo.Customer) {
	replyCard(ctx, msgId, CustomerInfoCardTemplate(customer))
}

func sendRoomInfoCard(ctx context.Context, msgId *string, customer bo.RoomBill) {
	replyCard(ctx, msgId, RoomInfoCardTemplate(customer))
}

func replyMsg(ctx context.Context, msg string, msgId string) {
	msg, i := processMessage(msg)
	if i != nil {
		return
	}
	client := initialization.GetLarkClient()
	content := larkim.NewTextMsgBuilder().
		Text(msg).
		Build()

	resp, err := client.Im.Message.Reply(ctx, larkim.NewReplyMessageReqBuilder().
		MessageId(msgId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeText).
			Uuid(uuid.New().String()).
			Content(content).
			Build()).
		Build())

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}
}

func replyCard(ctx context.Context, msgId *string, cardContent string,
) {
	client := initialization.GetLarkClient()
	resp, err := client.Im.Message.Reply(ctx, larkim.NewReplyMessageReqBuilder().
		MessageId(*msgId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeInteractive).
			Uuid(uuid.New().String()).
			Content(cardContent).
			Build()).
		Build())

	// 处理错误
	if err != nil {
		glog.Error(ctx, err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		glog.Errorf(ctx, "resp.Code:%v, resp.Msg:%s, resp.RequestId:%s", resp.Code, resp.Msg, resp.RequestId())
		return
	}
}

func MessagesCard(ctx context.Context,
	receiveId string,
	cardContent string,
) error {
	client := initialization.GetLarkClient()
	resp, err := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("user_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType(larkim.MsgTypeInteractive).
			Uuid(uuid.New().String()).
			Content(cardContent).
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

// withSplitLine 用于生成分割线
func withSplitLine() larkcard.MessageCardElement {
	splitLine := larkcard.NewMessageCardHr().
		Build()
	return splitLine
}

// withHeader 用于生成消息头
func withHeader(title string, color string) *larkcard.
	MessageCardHeader {
	if title == "" {
		title = "🤖️机器人提醒"
	}
	header := larkcard.NewMessageCardHeader().
		Template(color).
		Title(larkcard.NewMessageCardPlainText().
			Content(title).
			Build()).
		Build()
	return header
}

// withNote 用于生成纯文本脚注
func withNote(note string) larkcard.MessageCardElement {
	noteElement := larkcard.NewMessageCardNote().
		Elements([]larkcard.MessageCardNoteElement{larkcard.NewMessageCardPlainText().
			Content(note).
			Build()}).
		Build()
	return noteElement
}

// withMainMd 用于生成markdown消息体
func withMainMd(msg string) larkcard.MessageCardElement {
	msg, i := processMessage(msg)
	msg = processNewLine(msg)
	if i != nil {
		return nil
	}
	mainElement := larkcard.NewMessageCardDiv().
		Fields([]*larkcard.MessageCardField{larkcard.NewMessageCardField().
			Text(larkcard.NewMessageCardLarkMd().
				Content(msg).
				Build()).
			IsShort(true).
			Build()}).
		Build()
	return mainElement
}
