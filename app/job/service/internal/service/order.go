package service

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/job/service/internal/biz"
)

// OrderService is a order service.
type OrderService struct {
	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewOrderService new a greeter service.
func NewOrderService(uc *biz.GreeterUsecase, logger log.Logger) *OrderService {
	fmt.Println("new order service")
	return &OrderService{uc: uc, log: log.NewHelper(logger)}
}

// Start implements Start.Start
func (s *OrderService) Start() {
	s.log.Info("this is job")
	// return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
