package bill

import (
	"context"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
)

func init() {
	service.RegisterBill(New())
}

func New() *sBill {
	return &sBill{}
}

type sBill struct{}

func (s *sBill) Query(ctx context.Context, billQue do.Bill) ([]entity.Bill, error) {
	var bills []entity.Bill
	err := dao.Bill.Ctx(ctx).
		Where(billQue).
		Scan(&bills)
	if err != nil {
		return nil, err
	}
	return bills, nil
}

func (s *sBill) QueryPayment(ctx context.Context, billQue do.Bill) ([]entity.Bill, error) {
	var bills []entity.Bill
	err := dao.Bill.Ctx(ctx).
		Where(billQue).
		WhereGT(dao.Bill.Columns().Amount, 0).
		OrderDesc(dao.Bill.Columns().PayTime).
		Scan(&bills)
	if err != nil {
		return nil, err
	}
	return bills, nil
}

func (s *sBill) ListByIds(ctx context.Context, billIds *[]string) ([]entity.Bill, error) {
	var bills []entity.Bill
	err := dao.Bill.Ctx(ctx).WhereIn("id", *billIds).Scan(&bills)
	if err != nil {
		return nil, err
	}
	return bills, nil
}

func (s *sBill) BatchSave(ctx context.Context, bills *[]do.Bill) (err error) {
	_, err = dao.Bill.Ctx(ctx).Data(bills).Batch(10).Save()
	return
}

func (s *sBill) Create(ctx context.Context, bill entity.Bill) (err error) {
	_, err = dao.Bill.Ctx(ctx).Data(bill).Insert()
	return
}

func (s *sBill) Delete(ctx context.Context, billIds []string) (err error) {
	_, err = dao.Bill.Ctx(ctx).WhereIn(dao.Bill.Columns().Id, billIds).Delete()
	return
}
