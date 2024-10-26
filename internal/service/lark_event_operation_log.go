// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"lark-hotel-server/internal/model/do"
)

type (
	ILarkEventOperationLog interface {
		Create(ctx context.Context, log do.LarkEventOperationLog) (err error)
	}
)

var (
	localLarkEventOperationLog ILarkEventOperationLog
)

func LarkEventOperationLog() ILarkEventOperationLog {
	if localLarkEventOperationLog == nil {
		panic("implement not found for interface ILarkEventOperationLog, forgot register?")
	}
	return localLarkEventOperationLog
}

func RegisterLarkEventOperationLog(i ILarkEventOperationLog) {
	localLarkEventOperationLog = i
}
