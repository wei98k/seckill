package service

import (
	"context"

	pb "github.com/peter-wow/seckill/api/goods/service/v1"
	"github.com/peter-wow/seckill/app/goods/service/internal/biz"
)

func (s *GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	s.log.Info("GetGoods request: ", rv, req)
	return &pb.GetGoodsReply{
		Title: rv.Title,
		Intro: rv.Intro,
	}, err
}

func (s *GoodsService) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {
	rv, err := s.oc.GetOrders(ctx, req.Id)

	s.log.Info("request: ", rv, req)
	return &pb.GetOrdersReply{
		Sn: rv.Sn,
	}, err
}

func (s *GoodsService) CreateOrders(ctx context.Context, req *pb.CreateOrdersRequest) (*pb.CreateOrdersReply, error) {

	err := s.oc.CreateOrders(ctx, biz.Orders{
		Sn: req.Sn,
	})

	s.log.Info("request: ", req)

	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &pb.CreateOrdersReply{}, nil
}
