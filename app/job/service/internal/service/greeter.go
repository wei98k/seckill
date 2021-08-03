package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/peter-wow/seckill/api/job/service/v1"
	"github.com/peter-wow/seckill/app/job/service/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedJobServer

	uc  *biz.GreeterUsecase
	oc *biz.OrderQueueUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, oc *biz.OrderQueueUsecase, logger log.Logger) *GreeterService {
	fmt.Println("new greeter service")
	return &GreeterService{uc: uc, oc: oc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	m := &biz.SeckillOrder{
		OrderId: 111,
	}
	s.oc.Create(ctx, m)

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
