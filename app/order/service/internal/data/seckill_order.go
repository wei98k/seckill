package data

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
	"github.com/peter-wow/seckill/app/order/service/internal/data/kafka"
	"github.com/peter-wow/seckill/pkg/cache"
	"strconv"
)

var _ biz.SeckillOrderRepo = (*seckillOrderRepo)(nil)

type seckillOrderRepo struct {
	data *Data
	log *log.Helper
}

func (s seckillOrderRepo) SendKafkaOrder(ctx context.Context, order *biz.SeckillOrder) error {

	//雪花算法-id生成
	//node, err := snowflake.NewNode(1)
	//if err != nil {
	//	s.log.Error("snowflake generate error")
	//	return nil
	//}
	//id := node.Generate()

	// 订单消息队列处理
	msg := kafka.NewMessage("test-order", []byte("order"), map[string]string{
		"uid": strconv.FormatInt(order.UserId, 10),
		"goods_id": strconv.FormatInt(order.GoodsId, 10),
		"order_id": fmt.Sprintf("%s", "99882"),
	})

	return s.data.kafka.Send(ctx, msg)
}

func (s seckillOrderRepo) GetSeckillOrder(ctx context.Context, id int64) (*biz.SeckillOrder, error) {

	entityOrder, err := s.data.db.SeckillOrder.Get(ctx, id)

	return &biz.SeckillOrder{
		GoodsId: entityOrder.GoodsID,
		OrderId: entityOrder.OrderID,
		UserId: entityOrder.UserID,
	}, err
}

func (s seckillOrderRepo) PostSeckillOrder(ctx context.Context, seckillOrder *biz.SeckillOrder) error {
	entityOrder, err := s.data.db.SeckillOrder.Create().
							SetOrderID(seckillOrder.OrderId).
							SetGoodsID(seckillOrder.GoodsId).
							SetUserID(seckillOrder.UserId).Save(ctx)

	s.log.Debug(entityOrder, err)

	return err
}

func (s seckillOrderRepo) CreateSeckillOrder(ctx context.Context, seckillOrder *biz.SeckillOrder) error {

	//=====lock-分布式锁-加锁
	//var lockKey = "counter_lock"
	//var counterKey = "counter"
	//
	//// lock
	//resp := s.data.rdb.SetNX(ctx, lockKey, 1, time.Second*5)
	//lockSuccess, err := resp.Result()
	//
	//if err != nil || !lockSuccess {
	//	fmt.Println(err, "lock result: ", lockSuccess)
	//	return nil
	//}
	//
	//// counter ++
	//getResp := s.data.rdb.Get(ctx, counterKey)
	//cntValue, err := getResp.Int64()
	//if err == nil || err == redis.Nil {
	//	cntValue++
	//	resp := s.data.rdb.Set(ctx, counterKey, cntValue, 0)
	//	_, err := resp.Result()
	//	if err != nil {
	//		// log err
	//		println("set value error!")
	//	}
	//}
	//println("current counter is ", cntValue)
	//锁住之后处理的业务


	stock := s.data.rdb.Decr(ctx, cache.SeckillGoodStockKey(seckillOrder.GoodsId)).Val()

	if stock <= -1 {
		//商品库存已不足
		s.data.rdb.Del(ctx, cache.SeckillGoodOverKey(seckillOrder.GoodsId))
		return errors.Errorf(-1, "商品库存不足", "")
	}

	// 雪花ID生成
	node, err := snowflake.NewNode(1)
	if err != nil {
		s.log.Error("snowflake generate error")
		return nil
	}
	id := node.Generate()

	// 订单消息队列处理
	msg := kafka.NewMessage("seckill-order", []byte("order"), map[string]string{
		"uid": strconv.FormatInt(seckillOrder.UserId, 10),
		"goods_id": strconv.FormatInt(seckillOrder.GoodsId, 10),
		"order_id": fmt.Sprintf("%s", id),
	})

	s.data.kafka.Send(ctx, msg)
	//生成订单、减掉库存、(如何用户5分钟内未付款、再回执加回库存)


	//=====lock-分布式锁-释放锁
	//delResp := s.data.rdb.Del(ctx, lockKey)
	//unlockSuccess, err := delResp.Result()
	//if err == nil && unlockSuccess > 0 {
	//	println("unlock success!")
	//} else {
	//	println("unlock failed", err)
	//}

	return nil
}

func NewSeckillOrderRepo(data *Data, logger log.Logger) biz.SeckillOrderRepo {

	return &seckillOrderRepo{
		data: data,
		log: log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}
