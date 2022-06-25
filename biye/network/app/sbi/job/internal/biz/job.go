package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type JobRepo interface {
	IntegratingJob(ctx context.Context) error
}

type JobUseCase struct {
	repo JobRepo
	log  *log.Helper
}

func NewJobUseCase(repo JobRepo, logger log.Logger) *JobUseCase {
	return &JobUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/job")),
	}
}

func (j *JobUseCase) IntegratingJob(ctx context.Context) error {
	return j.repo.IntegratingJob(ctx)
}
