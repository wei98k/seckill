package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/helloMJW/seckill/api/order/service/v1"
	"github.com/helloMJW/seckill/app/order/service/internal/biz"
)



// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	pb.UnimplementedOrderServer

	uc *biz.OrderUsecase
	log *log.Helper
}

func NewOrderService(uc *biz.OrderUsecase, logger log.Logger) *OrderService  {
	return &OrderService{
		uc: uc,
		log: log.NewHelper(logger),
	}
}
