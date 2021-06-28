package service

import (
	"context"
	pb "github.com/helloMJW/seckill/api/order/service/v1"
	"github.com/helloMJW/seckill/app/order/service/internal/biz"
)

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	data := &biz.Order{
		Gid: req.Gid,
		Amount: 1,
	}

	err := s.uc.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderReply{},nil
}