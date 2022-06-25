package biz

import (
	"context"
	v1 "network/api/sbi/interface/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type InterfaceRepo interface {
	GetOrdersAndIntegrating(ctx context.Context, req *v1.GetOrdersAndIntegratingReq) (*v1.GetOrdersAndIntegratingReply, error)
}

type InterfaceUseCase struct {
	repo InterfaceRepo
	log  *log.Helper
}

func NewInterfaceUseCase(repo InterfaceRepo, logger log.Logger) *InterfaceUseCase {
	return &InterfaceUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
	}
}

func (s *InterfaceUseCase) GetOrdersAndIntegrating(ctx context.Context, req *v1.GetOrdersAndIntegratingReq) (*v1.GetOrdersAndIntegratingReply, error) {
	return s.repo.GetOrdersAndIntegrating(ctx, req)
}
