package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type OrderQueueRepo interface {
	CreateOrder(context.Context) error
}

type OrderQueueUsecase struct {
	repo OrderQueueRepo
	log *log.Helper
}

func NewOrderQueueUsecase(repo OrderQueueRepo, logger log.Logger) *OrderQueueUsecase {
	return &OrderQueueUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderQueueUsecase) CreateQueue(ctx context.Context) error {

	return uc.repo.CreateOrder(ctx)
}
