package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

//Domain Object
type Sbi struct {
	Id       int64
	Title    string
	Artist   string
	Price    float64
	CreateAt time.Time
}

// Domain Object
type Orders struct {
	Id       int64
	UserId   int64
	SbiId    int64
	Price    float64
	Receiver string
	Address  string
	Mobile   string
}

type ListSbi struct {
	Count int64
	Sbis  []*Sbi
}

type SbiRepo interface {
	ListSbi(ctx context.Context, pageNum, pageSiz int64) (ListSbi, error)
	GetSbiById(ctx context.Context, id int64) (*Sbi, error)
	CreateOrders(ctx context.Context, orders *Orders) error
}

type SbiUseCase struct {
	repo SbiRepo
	log  *log.Helper
}

func NewSbiUseCase(repo SbiRepo, logger log.Logger) *SbiUseCase {
	return &SbiUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/sbi")),
	}
}

func (uc *SbiUseCase) ListSbi(ctx context.Context, pageNum, pageSize int64) (ListSbi, error) {
	return uc.repo.ListSbi(ctx, pageNum, pageSize)
}

func (uc *SbiUseCase) GetSbiById(ctx context.Context, id int64) (*Sbi, error) {
	return uc.repo.GetSbiById(ctx, id)
}

func (uc *SbiUseCase) CreateOrders(ctx context.Context, orders *Orders) error {
	return uc.repo.CreateOrders(ctx, orders)
}
