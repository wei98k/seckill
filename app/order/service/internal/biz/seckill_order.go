package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SeckillOrder struct {
	UserId int64
	OrderId int64
	GoodsId int64
}

type SeckillOrderRepo interface {
	CreateSeckillOrder(ctx context.Context, seckillOrder *SeckillOrder) error
}

type SeckillOrderUsecase struct {
	repo SeckillOrderRepo
	log *log.Helper
}

func NewSeckillOrderUsecase(repo SeckillOrderRepo, logger log.Logger) *SeckillOrderUsecase {
	return &SeckillOrderUsecase{
		repo: repo,
		log: log.NewHelper(logger),
	}
}

func (uc *SeckillOrderUsecase) CreateSeckillOrder(ctx context.Context, o *SeckillOrder) error {

	return uc.repo.CreateSeckillOrder(ctx, o)
}

