package service

import (
	"context"
	"fmt"

	"github.com/peter-wow/seckill/app/job/service/internal/biz"
)

// OrderService is a order service.
type OrderService struct {
	oc *biz.OrderQueueUsecase
}

func (s *OrderService) Create(ctx context.Context, m *biz.SeckillOrder) {
	fmt.Println("service - create")
	// 读取kafka消息后、准备写入数据库的时候、数据库崩溃了或其他意外情况导致无法写入数据、那这个数据只能丢失了

	//TODO::

	//分布式-事务开始-start

	//m := &biz.SeckillOrder{
	//	OrderId: 111,
	//}

	//TODO::子事务、插入秒杀订单表

	//TODO::子事务、插入订单表

	//TODO::子事务、插入订单详细表

	//分布式-事务结束-end

	s.oc.Create(ctx, m)

}

// NewOrderService new a greeter service.
func NewOrderService(oc *biz.OrderQueueUsecase) *OrderService {
	return &OrderService{oc: oc}
}
