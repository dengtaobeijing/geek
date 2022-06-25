package service

import (
	"context"
	"time"

	v1 "network/api/sbi/service/v1"
	"network/app/sbi/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type SbiService struct {
	v1.UnimplementedSbiServer

	ac  *biz.SbiUseCase
	log *log.Helper
}

func NewSbiService(ac *biz.SbiUseCase, logger log.Logger) *SbiService {
	return &SbiService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/sbi")),
	}

}

func (s *SbiService) ListSbi(ctx context.Context, req *v1.ListSbiReq) (*v1.ListSbiReply, error) {
	rv, err := s.ac.ListSbi(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	sbis := make([]*v1.ListSbiReply_Sbi, 0)
	for _, r := range rv.Sbis {
		sbis = append(sbis, &v1.ListSbiReply_Sbi{
			Id:       r.Id,
			Title:    r.Title,
			Artist:   r.Artist,
			Price:    r.Price,
			CreateAt: time.Unix(r.CreateAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.ListSbiReply{
		Count: rv.Count,
		Sbis:  sbis,
	}, nil
}

func (s *SbiService) GetSbiById(ctx context.Context, req *v1.GetSbiByIdReq) (*v1.GetSbiByIdReply, error) {
	rv, err := s.ac.GetSbiById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	sbi := &v1.GetSbiByIdReply{
		Sbi: &v1.GetSbiByIdReply_Sbi{
			Id:       rv.Id,
			Title:    rv.Title,
			Artist:   rv.Artist,
			Price:    rv.Price,
			CreateAt: time.Unix(rv.CreateAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		},
	}
	return sbi, nil
}

func (s *SbiService) CreateOrders(ctx context.Context, req *v1.CreateOrdersReq) (*v1.CreateOrdersReply, error) {
	order := biz.Orders{
		UserId:   req.Orders.UserId,
		SbiId:    req.Orders.SbiId,
		Price:    req.Orders.Price,
		Receiver: req.Orders.Receiver,
		Address:  req.Orders.Address,
		Mobile:   req.Orders.Mobile,
	}
	return &v1.CreateOrdersReply{}, s.ac.CreateOrders(ctx, &order)
}
