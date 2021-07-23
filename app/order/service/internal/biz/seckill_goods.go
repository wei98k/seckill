package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// SeckillGoods 定义对应数据库实体
type SeckillGoods struct {
	GoodsId int64
	SeckillPrice float64
	StockCount int64
	StartDate time.Time
	EndDate time.Time
}
// 定义方法
type SeckillGoodsRepo interface {
	GetGoods(ctx context.Context,id int64) (*SeckillGoods, error)
	DecrGoodsStock(ctx context.Context,id int64) error
}

type SeckillGoodsUsecase struct {
	repo SeckillGoodsRepo
	log *log.Helper
}

func NewSeckillGoodsUsecase (repo SeckillGoodsRepo, logger log.Logger) *SeckillGoodsUsecase {
	return &SeckillGoodsUsecase {
		repo: repo,
		log: log.NewHelper(logger),
	}
}

func (g SeckillGoodsUsecase) GetSeckillGoods(ctx context.Context, id int64) (*SeckillGoods, error) {
	goods, err := g.repo.GetGoods(ctx, id)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (g SeckillGoodsUsecase) DecrGoodsStock(ctx context.Context, id int64) error {
	err := g.repo.DecrGoodsStock(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
