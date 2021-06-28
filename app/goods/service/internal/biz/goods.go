package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Goods struct {
	Id int64
	Title string
}

type GoodsUsecase struct {
	repo GoodsRepo
	log *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{
		repo: repo,
		log: log.NewHelper(logger),
	}
}

type GoodsRepo interface {
	GetGoods(ctx context.Context, id int64) (*Goods, error)
}

func (uc *GoodsUsecase) Get(ctx context.Context, id int64) (*Goods, error) {
	return uc.repo.GetGoods(ctx, id)
}