package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)

var _ biz.SeckillOrderRepo = (*seckillOrderRepo)(nil)

type seckillOrderRepo struct {
	data *Data
	log *log.Helper
}

func (s seckillOrderRepo) CreateSeckillOrder(ctx context.Context, seckillOrder *biz.SeckillOrder) error {

	//判断库存剩余量
	stock, err := s.data.rdb.Decr(ctx, "SECKILL:ORDER:ID:12").Result()

	if stock < 0 {
		//TODO::通知前端把按钮改成卖完了
		//TODO::接口不在接收请求
		s.log.Infof("剩余stock: %v, error: %v", stock, err)
		return errors.Errorf(-1, "卖完了。。。", "")
	}

	res, err := s.data.db.SeckillOrder.Create().
					SetGoodsID(seckillOrder.GoodsId).
					SetOrderID(seckillOrder.OrderId).
					SetUserID(seckillOrder.UserId).Save(ctx)
	if err != nil {
		return err
	}
	s.log.Info("seckill-data: res ", res)
	return nil
}

func NewSeckillOrderRepo(data *Data, logger log.Logger) biz.SeckillOrderRepo {
	return &seckillOrderRepo{
		data: data,
		log: log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}
