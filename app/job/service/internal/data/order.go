package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/job/service/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

type SeckillOrder struct {
	UserId int64
	OrderId int64
	GoodsId int64
}

func (o *orderRepo) CreateOrder(ctx context.Context, oo *biz.SeckillOrder) (*biz.SeckillOrder, error) {

	m := SeckillOrder{
		UserId: oo.UserId,
		OrderId: oo.OrderId,
		GoodsId: oo.GoodsId,
	}

	result := o.data.db.WithContext(ctx).Create(m)

	return &biz.SeckillOrder{
		OrderId: oo.OrderId,
	}, result.Error
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderQueueRepo {
	return &orderRepo {
		data: data,
		log: log.NewHelper(logger),
	}
}