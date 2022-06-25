package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TaskRepo interface {
	IntegratingCount(ctx context.Context) error
	GetIntegrating(ctx context.Context, userId int64) (int64, error)
}

type TaskUseCase struct {
	repo TaskRepo
	log  *log.Helper
}

func NewTaskUseCase(repo TaskRepo, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/task")),
	}
}

func (j *TaskUseCase) IntegratingCount(ctx context.Context) error {
	return j.repo.IntegratingCount(ctx)
}

func (j *TaskUseCase) GetIntegrating(ctx context.Context, userId int64) (int64, error) {
	return j.repo.GetIntegrating(ctx, userId)
}
