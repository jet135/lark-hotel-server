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
	ICustomer interface {
		FindByNameIn(ctx context.Context, names []string) (*[]entity.Customer, error)
		FindByPhoneIn(ctx context.Context, phones []string) (*[]entity.Customer, error)
		Query(ctx context.Context, customer do.Customer) (*[]entity.Customer, error)
		BatchCreate(ctx context.Context, customers *[]do.Customer) (err error)
		Create(ctx context.Context, bill do.Customer) (err error)
	}
)

var (
	localCustomer ICustomer
)

func Customer() ICustomer {
	if localCustomer == nil {
		panic("implement not found for interface ICustomer, forgot register?")
	}
	return localCustomer
}

func RegisterCustomer(i ICustomer) {
	localCustomer = i
}
