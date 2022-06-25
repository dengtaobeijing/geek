package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

//Domain Object
type Orders struct {
	Id       int64
	UserId   int64
	SbiId    int64
	Price    float64
	Receiver string
	Address  string
	Mobile   string
	CreateAt time.Time
}

type OrdersRepo interface {
	UpdateOrders(ctx context.Context, userId int64, b *Orders) error
}

type OrdersUseCase struct {
	repo OrdersRepo
	log  *log.Helper
}

func NewOrdersUseCase(repo OrdersRepo, logger log.Logger) *OrdersUseCase {
	return &OrdersUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/orders")),
	}
}

func (uc *OrdersUseCase) UpdateOrders(ctx context.Context, userId int64, b *Orders) error {
	return uc.repo.UpdateOrders(ctx, userId, b)
}
