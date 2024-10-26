package bill

import (
	"context"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
)

func init() {
	service.RegisterCustomer(New())
}

func New() *sCustomer {
	return &sCustomer{}
}

type sCustomer struct{}

func (s *sCustomer) FindByNameIn(ctx context.Context, names []string) (*[]entity.Customer, error) {
	var customers []entity.Customer
	err := dao.Customer.Ctx(ctx).WhereIn(dao.Customer.Columns().Name, names).Scan(&customers)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (s *sCustomer) FindByPhoneIn(ctx context.Context, phones []string) (*[]entity.Customer, error) {
	var customers []entity.Customer
	err := dao.Customer.Ctx(ctx).WhereIn(dao.Customer.Columns().Phone, phones).Scan(&customers)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (s *sCustomer) Query(ctx context.Context, customer do.Customer) (*[]entity.Customer, error) {
	var customers []entity.Customer
	err := dao.Customer.Ctx(ctx).Where(customer).Scan(&customers)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}
func (s *sCustomer) BatchCreate(ctx context.Context, customers *[]do.Customer) (err error) {
	_, err = dao.Customer.Ctx(ctx).Data(customers).Batch(10).Insert()
	return
}

func (s *sCustomer) Create(ctx context.Context, bill do.Customer) (err error) {
	_, err = dao.Customer.Ctx(ctx).Data(bill).Insert()
	return
}
