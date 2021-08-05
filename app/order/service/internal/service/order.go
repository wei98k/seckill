package service

import (
	"context"
	pb "github.com/peter-wow/seckill/api/order/service/v1"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {

	s.log.Infof("CreateOrder request: %v", req)

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

func (s *OrderService) CreateSeckillOrder(ctx context.Context, req *pb.CreateSeckillOrderRequest) (*pb.CreateSeckillOrderReply, error) {

	s.log.Info("input data %v ", req)

	//验证秒杀活动
	goods, err1 := s.goods.GetSeckillGoods(ctx, req.Gid)

	s.goods.DecrGoodsStock(ctx, req.Gid)

	if err1 != nil {
		s.log.Errorf("get goods error %v", err1)
		return nil, err1
	}
	s.log.Infof("get goods data: %v", goods)

	param := &biz.SeckillOrder{
		UserId: 88,
		GoodsId: req.Gid,
		OrderId: 99,
	}

	err := s.so.CreateSeckillOrder(ctx, param)

	s.log.Info("seckill-order-service: err ", err)

	return &pb.CreateSeckillOrderReply{},nil
}