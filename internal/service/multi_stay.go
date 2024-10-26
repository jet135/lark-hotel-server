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
	IMultiStay interface {
		Create(ctx context.Context, MultiStays *[]do.MultiStay) (err error)
		SettingDoneByRoomNumberAndCheckInTime(ctx context.Context, roomNumber string, checkInTime time.Time) (err error)
		SettingDoneByBillIdIn(ctx context.Context, billIds []string) (err error)
		SettingDone(ctx context.Context, id int64) (err error)
		DeleteByBillIds(ctx context.Context, billIds []string) (err error)
		FindByDone(ctx context.Context, done int8) ([]entity.MultiStay, error)
		FindByBillIdIn(ctx context.Context, billIds []string) ([]entity.MultiStay, error)
	}
)

var (
	localMultiStay IMultiStay
)

func MultiStay() IMultiStay {
	if localMultiStay == nil {
		panic("implement not found for interface IMultiStay, forgot register?")
	}
	return localMultiStay
}

func RegisterMultiStay(i IMultiStay) {
	localMultiStay = i
}
