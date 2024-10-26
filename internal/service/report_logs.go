// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"lark-hotel-server/api/report"
	"lark-hotel-server/internal/model/entity"
)

type (
	IReportLogs interface {
		Process(ctx context.Context, req report.Req) (err error)
		Create(ctx context.Context, reportLog *entity.ReportLogs) (err error)
	}
)

var (
	localReportLogs IReportLogs
)

func ReportLogs() IReportLogs {
	if localReportLogs == nil {
		panic("implement not found for interface IReportLogs, forgot register?")
	}
	return localReportLogs
}

func RegisterReportLogs(i IReportLogs) {
	localReportLogs = i
}
