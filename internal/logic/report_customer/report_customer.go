package report_customer

import (
	"context"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"time"
)

func init() {
	service.RegisterReportCustomer(New())
}

func New() *sReportCustomer {
	return &sReportCustomer{}
}

type sReportCustomer struct{}

func (s *sReportCustomer) Create(ctx context.Context, reportCustomer do.ReportCustomer) (err error) {
	_, err = dao.ReportCustomer.Ctx(ctx).Data(reportCustomer).Insert()
	return
}

func (s *sReportCustomer) ListByGTECreateAt(ctx context.Context, createAt time.Time) (*[]entity.ReportCustomer, error) {
	var reportCustomers []entity.ReportCustomer
	err := dao.ReportCustomer.Ctx(ctx).WhereGTE(dao.ReportCustomer.Columns().CreatedAt, createAt).Scan(&reportCustomers)
	if err != nil {
		return nil, err
	}
	return &reportCustomers, nil
}
