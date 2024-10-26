// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
)

type (
	IReportCustomer interface {
		Create(ctx context.Context, reportCustomer do.ReportCustomer) (err error)
		ListByGTECreateAt(ctx context.Context, createAt time.Time) (*[]entity.ReportCustomer, error)
	}
)

var (
	localReportCustomer IReportCustomer
)

func ReportCustomer() IReportCustomer {
	if localReportCustomer == nil {
		panic("implement not found for interface IReportCustomer, forgot register?")
	}
	return localReportCustomer
}

func RegisterReportCustomer(i IReportCustomer) {
	localReportCustomer = i
}
