package service

import (
	"github.com/peter-wow/seckill/app/job/service/internal/biz"
)

// OrderService is a order service.
type OrderService struct {
	uc  *biz.GreeterUsecase
	oc *biz.OrderQueueUsecase
}

// NewOrderService new a greeter service.
func NewOrderService(uc *biz.GreeterUsecase, oc *biz.OrderQueueUsecase) *OrderService {

	return &OrderService{uc: uc, oc: oc}
}
