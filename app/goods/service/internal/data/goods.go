package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/helloMJW/seckill/app/goods/service/internal/biz"
)

var _ biz.GoodsRepo = (*goodsRepo)(nil)

type goodsRepo struct {
	data *Data
	log *log.Helper
}

func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log: log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (g goodsRepo) GetGoods(ctx context.Context, id int64) (*biz.Goods, error) {
	panic("implement me")
}