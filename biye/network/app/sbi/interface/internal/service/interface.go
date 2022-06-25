package service

import (
	"context"
	v1 "network/api/sbi/interface/v1"
	"network/app/sbi/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type InterfaceService struct {
	v1.UnimplementedInterfaceServer

	ac  *biz.InterfaceUseCase
	log *log.Helper
}

func NewInterfaceService(ac *biz.InterfaceUseCase, logger log.Logger) *InterfaceService {
	return &InterfaceService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
	}

}

func (s *InterfaceService) GetOrdersAndIntegrating(ctx context.Context, req *v1.GetOrdersAndIntegratingReq) (*v1.GetOrdersAndIntegratingReply, error) {
	return s.ac.GetOrdersAndIntegrating(ctx, req)
}
