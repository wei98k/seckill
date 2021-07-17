package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)

var _ biz.SeckillGoodsRepo = (*seckillGoodsRepo)(nil)

type seckillGoodsRepo struct {
	data *Data
	log *log.Helper
}

func (s seckillGoodsRepo) GetGoods(ctx context.Context, id int64) (*biz.SeckillGoods, error) {

	//两种查询方式
	// 1. 先通过ID获取出整条数据 然后判断是否开始、结束、库存是否充足
	// 2. 通过where加入活动时间、库存条件 如果没有数据表示活动未开始或结束、或已售空
	s.log.Infof("GetGoods param: %v", id)
	g, err := s.data.db.SeckillGoods.Get(ctx, id)
	s.log.Infof("GetGoods result: %v %v", g, err)
	if err != nil {
		return nil, err
	}

	//TODO::判断活动是否开始

	//TODO::判断活动是否结束

	//TODO::判断库存是否充足

	return &biz.SeckillGoods{
		GoodsId: g.GoodsID,
		SeckillPrice: g.SeckillPrice,
		StockCount: g.StockCount,
	}, nil

}

func NewSeckillGoodsRepo (data *Data, logger log.Logger) biz.SeckillGoodsRepo {
	return &seckillGoodsRepo{
		data: data,
		log: log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}
