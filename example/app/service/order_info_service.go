package service

import (
	"github.com/lrayt/sparrow/core/runtime"
	"github.com/lrayt/sparrow/example/app/dao"
	"github.com/lrayt/sparrow/example/app/model"
	"github.com/lrayt/sparrow/example/app/protobuf/pb"
	"github.com/lrayt/sparrow/pkg/date"
	"github.com/lrayt/sparrow/pkg/uuid"
	"github.com/lrayt/sparrow/ts_error"
)

type OrderInfoService struct {
	orderInfoDao *dao.OrderInfoDao
}

func NewOrderInfoService(orderInfoDao *dao.OrderInfoDao) *OrderInfoService {
	return &OrderInfoService{orderInfoDao: orderInfoDao}
}

func (s OrderInfoService) CreateOrder(ctx *runtime.Context, req *pb.OrderCreateRequest) *pb.OrderCreateResponse {
	if err := s.orderInfoDao.SaveOrderInfo(&model.OrderInfo{
		OrderNO:      uuid.GenUUID(),
		ProductId:    req.ProductId,
		Status:       model.OrderStatusUnpaid,
		ProductCount: req.Count,
		AmountTotal:  10,
		LogisticsFee: 10,
		AddressId:    req.AddressId,
		UserId:       ctx.UserId(),
		PayDate:      date.NowTime(),
	}); err != nil {
		return &pb.OrderCreateResponse{
			Msg:  ctx.Log().NewErrorF("创建失败", err).Error(),
			Code: ts_error.DataToDBErr,
		}
	}
	return &pb.OrderCreateResponse{
		Msg:  ctx.Log().Success("创建成功"),
		Code: ts_error.Success,
	}
}
