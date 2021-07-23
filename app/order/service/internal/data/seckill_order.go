package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
	"time"
)

var _ biz.SeckillOrderRepo = (*seckillOrderRepo)(nil)

type seckillOrderRepo struct {
	data *Data
	log *log.Helper
}

func (s seckillOrderRepo) CreateSeckillOrder(ctx context.Context, seckillOrder *biz.SeckillOrder) error {

	//判断库存剩余量
	//stock, err := s.data.rdb.Decr(ctx, "SECKILL:ORDER:ID:12").Result()
	//
	//if stock < 0 {
	//	//TODO::通知前端把按钮改成卖完了
	//	//TODO::接口不在接收请求
	//	s.log.Infof("剩余stock: %v, error: %v", stock, err)
	//	return errors.Errorf(-1, "卖完了。。。", "")
	//}

	// 分布式锁-redis

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := s.data.rdb.SetNX(ctx, lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return nil
	}

	// counter ++
	getResp := s.data.rdb.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := s.data.rdb.Set(ctx, counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	//锁住之后处理的业务
	res, err := s.data.db.SeckillOrder.Create().
					SetGoodsID(seckillOrder.GoodsId).
					SetOrderID(seckillOrder.OrderId).
					SetUserID(seckillOrder.UserId).Save(ctx)
	if err != nil {
		return err
	}
	s.log.Info("seckill-data: res ", res)


	//释放锁
	delResp := s.data.rdb.Del(ctx, lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}

	return nil
}

func NewSeckillOrderRepo(data *Data, logger log.Logger) biz.SeckillOrderRepo {
	return &seckillOrderRepo{
		data: data,
		log: log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}
