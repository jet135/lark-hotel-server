package lark_app_info

import (
	"context"
	"lark-hotel-server/internal/dao"
	"lark-hotel-server/internal/model/do"
	"lark-hotel-server/internal/model/entity"
	"lark-hotel-server/internal/service"
	"time"
)

func init() {
	service.RegisterLarkAppInfo(New())
}

func New() *sLarkAppInfo {
	return &sLarkAppInfo{}
}

type sLarkAppInfo struct{}

func (s *sLarkAppInfo) Query(ctx context.Context, info do.LarkAppInfo) ([]entity.LarkAppInfo, error) {
	var appInfos []entity.LarkAppInfo
	err := dao.LarkAppInfo.Ctx(ctx).Data(info).Scan(&appInfos)
	if err != nil {
		return nil, err
	}
	return appInfos, nil
}

func (s *sLarkAppInfo) GetByBelongDate(ctx context.Context, belongDate time.Time) (*entity.LarkAppInfo, error) {
	m := dao.LarkAppInfo.Ctx(ctx)
	record, err := m.Where(&do.LarkAppInfo{BelongDate: belongDate}).One()
	if err != nil {
		return nil, err
	}
	var larkAppInfo entity.LarkAppInfo
	if record.IsEmpty() {
		return nil, nil
	}
	err = record.Struct(&larkAppInfo)
	if err != nil {
		return nil, err
	}
	return &larkAppInfo, nil
}

func (s *sLarkAppInfo) Create(ctx context.Context, info *entity.LarkAppInfo) error {
	_, err := dao.LarkAppInfo.Ctx(ctx).Data(info).Insert()
	return err
}
