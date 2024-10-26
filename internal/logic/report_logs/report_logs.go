package report_logs

import (
	"context"
	"lark-hotel-server/api/report"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"lark-hotel-server/internal/third_service"
	"lark-hotel-server/internal/utils"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

func init() {
	service.RegisterReportLogs(New())
}

func New() *sReportLogs {
	return &sReportLogs{}
}

type sReportLogs struct{}

func (s *sReportLogs) Process(ctx context.Context, req report.Req) (err error) {
	var message report.Message
	if err = gconv.Struct(req.RequestBody, &message); err != nil {
		return
	}
	if message.ServiceDescript.ServiceName == "/submitSjrzr" {
		if err = s.Create(ctx, &entity.ReportLogs{
			OriginalUrl: req.OriginalUrl,
			Method:      req.Method,
			RequestBody: gconv.String(req.RequestBody),
		}); err != nil {
			return
		}
		var dataObj report.Data
		if err = gconv.Struct(message.ServiceDescript.Data, &dataObj); err != nil {
			return
		}
		dataObj.Date = time.Now()
		dataObj.Shift = utils.GetShift()
		// 上报数据发给飞书文档
		_, err = third_service.SendToLark(dataObj, ctx)
		if err != nil {
			g.Log().Errorf(ctx, "sendToLark err .. :%v", err)
		}
		// 保存数据库
		_ = service.ReportCustomer().Create(ctx, do.ReportCustomer{
			Id:         uuid.New().String(),
			Name:       dataObj.Name,
			BirthDate:  dataObj.BirthDate,
			Czdz:       dataObj.Czdz,
			IdCode:     dataObj.IdCode,
			IdType:     dataObj.IdType,
			Phone:      dataObj.Phone,
			XzqhTitle:  dataObj.XzqhTitle,
			RoomNumber: dataObj.RoomNumber,
		})
	}
	return nil
}

func (s *sReportLogs) Create(ctx context.Context, reportLog *entity.ReportLogs) (err error) {
	_, err = dao.ReportLogs.Ctx(ctx).Data(reportLog).Insert()
	return
}
