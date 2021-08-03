package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type OrderQueueRepo interface {
	CreateOrder(ctx context.Context, o *SeckillOrder) (*SeckillOrder, error)
}

type SeckillOrder struct {
	UserId int64
	OrderId int64
	GoodsId int64
}

type OrderQueueUsecase struct {
	repo OrderQueueRepo
	log *log.Helper
}

func NewOrderQueueUsecase(repo OrderQueueRepo, logger log.Logger) *OrderQueueUsecase {
	return &OrderQueueUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderQueueUsecase) Create(ctx context.Context, o *SeckillOrder) (*SeckillOrder, error) {
	return uc.repo.CreateOrder(ctx, o)
}
