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

type ListSbi struct {
	Count int64
	Sbis  []*Sbi
}

type SbiRepo interface {
	UpdateSbi(ctx context.Context, userId int64, sbi *Sbi) error
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

func (uc *SbiUseCase) UpdateSbi(ctx context.Context, userId int64, sbi *Sbi) error {
	return uc.repo.UpdateSbi(ctx, userId, sbi)
}
