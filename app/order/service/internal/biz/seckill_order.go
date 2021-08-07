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
	PostSeckillOrder(ctx context.Context, seckillOrder *SeckillOrder) error
	GetSeckillOrder(ctx context.Context, id int64) (*SeckillOrder, error)
	SendKafkaOrder(ctx context.Context, order *SeckillOrder) error
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

func (uc *SeckillOrderUsecase) PostOrder(ctx context.Context, o *SeckillOrder) error {
	return uc.repo.PostSeckillOrder(ctx, o)
}

func (uc *SeckillOrderUsecase) GetOrder(ctx context.Context, id int64) (*SeckillOrder, error) {
	return uc.repo.GetSeckillOrder(ctx, id)
}

func (uc *SeckillOrderUsecase) SendKafka(ctx context.Context, order *SeckillOrder) error {
	return uc.repo.SendKafkaOrder(ctx, order)
}


