// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
)

type (
	IBill interface {
		Query(ctx context.Context, billQue do.Bill) ([]entity.Bill, error)
		QueryPayment(ctx context.Context, billQue do.Bill) ([]entity.Bill, error)
		ListByIds(ctx context.Context, billIds *[]string) ([]entity.Bill, error)
		BatchSave(ctx context.Context, bills *[]do.Bill) (err error)
		Create(ctx context.Context, bill entity.Bill) (err error)
		Delete(ctx context.Context, billIds []string) (err error)
	}
)

var (
	localBill IBill
)

func Bill() IBill {
	if localBill == nil {
		panic("implement not found for interface IBill, forgot register?")
	}
	return localBill
}

func RegisterBill(i IBill) {
	localBill = i
}
