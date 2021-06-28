package service

import (
	"context"
	pb "github.com/helloMJW/seckill/api/goods/service/v1"
)

func (s *GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	s.log.Info("GetGoods request: ", rv, req)
	return &pb.GetGoodsReply{
		Title: rv.Title,
		Intro: rv.Intro,
	}, err
}