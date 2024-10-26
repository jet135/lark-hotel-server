package report_logs

import (
	"context"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/service"
)

func init() {
	service.RegisterLarkEventOperationLog(New())
}

func New() *sLarkEventOperationLog {
	return &sLarkEventOperationLog{}
}

type sLarkEventOperationLog struct{}

func (s *sLarkEventOperationLog) Create(ctx context.Context, log do.LarkEventOperationLog) (err error) {
	_, err = dao.LarkEventOperationLog.Ctx(ctx).Data(log).Insert()
	return
}
