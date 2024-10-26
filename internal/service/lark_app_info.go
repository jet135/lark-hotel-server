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
	ILarkAppInfo interface {
		Query(ctx context.Context, info do.LarkAppInfo) ([]entity.LarkAppInfo, error)
		GetByBelongDate(ctx context.Context, belongDate time.Time) (*entity.LarkAppInfo, error)
		Create(ctx context.Context, info *entity.LarkAppInfo) error
	}
)

var (
	localLarkAppInfo ILarkAppInfo
)

func LarkAppInfo() ILarkAppInfo {
	if localLarkAppInfo == nil {
		panic("implement not found for interface ILarkAppInfo, forgot register?")
	}
	return localLarkAppInfo
}

func RegisterLarkAppInfo(i ILarkAppInfo) {
	localLarkAppInfo = i
}
