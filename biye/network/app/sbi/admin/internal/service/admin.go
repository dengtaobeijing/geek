package service

import (
	"context"
	"time"

	v1 "network/api/sbi/admin/v1"
	"network/app/sbi/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	v1.UnimplementedAdminServer

	auc *biz.SbiUseCase
	ouc *biz.OrdersUseCase
	log *log.Helper
}

func NewAdminService(auc *biz.SbiUseCase, ouc *biz.OrdersUseCase, logger log.Logger) *AdminService {
	return &AdminService{
		auc: auc,
		ouc: ouc,
		log: log.NewHelper(log.With(logger, "module", "service/admin")),
	}
}

func (s *AdminService) UpdateSbi(ctx context.Context, req *v1.UpdateSbiReq) (*v1.UpdateSbiReply, error) {
	crtAt, err := time.ParseInLocation("2006-01-02 15:04:05", req.Sbi.CreateAt, time.Local)
	if err != nil {
		return &v1.UpdateSbiReply{}, err
	}
	sbi := biz.Sbi{
		Id:       req.Sbi.Id,
		Title:    req.Sbi.Title,
		Artist:   req.Sbi.Artist,
		Price:    req.Sbi.Price,
		CreateAt: crtAt,
	}
	return &v1.UpdateSbiReply{}, s.auc.UpdateSbi(ctx, req.UserId, &sbi)
}

func (s *AdminService) UpdateOrders(ctx context.Context, req *v1.UpdateOrdersReq) (*v1.UpdateOrdersReply, error) {
	crtAt, err := time.ParseInLocation("2006-01-02 15:04:05", req.Orders.CreateAt, time.Local)
	if err != nil {
		return &v1.UpdateOrdersReply{}, err
	}
	order := biz.Orders{
		Id:       req.Orders.Id,
		UserId:   req.Orders.UserId,
		SbiId:    req.Orders.SbiId,
		Price:    req.Orders.Price,
		Receiver: req.Orders.Receiver,
		Address:  req.Orders.Address,
		Mobile:   req.Orders.Mobile,
		CreateAt: crtAt,
	}
	return &v1.UpdateOrdersReply{}, s.ouc.UpdateOrders(ctx, req.Orders.UserId, &order)
}
