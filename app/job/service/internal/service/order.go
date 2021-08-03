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

	//m := &biz.SeckillOrder{
	//	OrderId: 111,
	//}
	s.oc.Create(ctx, m)

}

// NewOrderService new a greeter service.
func NewOrderService(oc *biz.OrderQueueUsecase) *OrderService {

	return &OrderService{ oc: oc}
}
