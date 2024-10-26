package multi_stay

import (
	"context"
	"lark-hotel-server/internal/consts"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"time"
)

func init() {
	service.RegisterMultiStay(New())
}

func New() *sMultiStay {
	return &sMultiStay{}
}

type sMultiStay struct{}

func (s *sMultiStay) Create(ctx context.Context, MultiStays *[]do.MultiStay) (err error) {
	_, err = dao.MultiStay.Ctx(ctx).Data(MultiStays).Batch(10).Insert()
	return
}

func (s *sMultiStay) SettingDoneByRoomNumberAndCheckInTime(ctx context.Context, roomNumber string, checkInTime time.Time) (err error) {
	_, err = dao.MultiStay.Ctx(ctx).Data(do.MultiStay{Done: consts.BusinessFlagYes}).
		Where(dao.MultiStay.Columns().RoomNumber, roomNumber).
		Where(dao.MultiStay.Columns().CheckinTime, checkInTime).
		Where(dao.MultiStay.Columns().Done, consts.BusinessFlagNo).
		Update()
	return
}

func (s *sMultiStay) SettingDoneByBillIdIn(ctx context.Context, billIds []string) (err error) {
	_, err = dao.MultiStay.Ctx(ctx).Data(do.MultiStay{Done: consts.BusinessFlagYes}).WhereIn(dao.MultiStay.Columns().BillId, billIds).Update()
	return
}

func (s *sMultiStay) SettingDone(ctx context.Context, id int64) (err error) {
	_, err = dao.MultiStay.Ctx(ctx).Data(do.MultiStay{Done: consts.BusinessFlagYes}).Where("id", id).Update()
	return
}

func (s *sMultiStay) DeleteByBillIds(ctx context.Context, billIds []string) (err error) {
	_, err = dao.MultiStay.Ctx(ctx).WhereIn(dao.MultiStay.Columns().BillId, billIds).Delete()
	return
}

func (s *sMultiStay) FindByDone(ctx context.Context, done int8) ([]entity.MultiStay, error) {
	var multiStays []entity.MultiStay
	err := dao.MultiStay.Ctx(ctx).Where("done", done).Scan(&multiStays)
	return multiStays, err
}

func (s *sMultiStay) FindByBillIdIn(ctx context.Context, billIds []string) ([]entity.MultiStay, error) {
	var multiStays []entity.MultiStay
	err := dao.MultiStay.Ctx(ctx).WhereIn(dao.MultiStay.Columns().BillId, billIds).Scan(&multiStays)
	return multiStays, err
}
