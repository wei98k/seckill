package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
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

// tcc-try
func (s *GoodsService) CreateOrdersTccTry(ctx context.Context, req *pb.CreateOrdersRequest) (*pb.CreateOrdersReply, error) {
	s.log.Info("ttc-try")
	return &pb.CreateOrdersReply{
		DtmResult: "SUCCESS",
	}, nil
}

// tcc-confirm
func (s *GoodsService) CreateOrdersConfirm(ctx context.Context, req *pb.CreateOrdersRequest) (*pb.CreateOrdersReply, error) {
	s.log.Info("ttc-confirm")
	if tr, ok := transport.FromServerContext(ctx); ok {
		// kind := tr.Kind().String()
		// operation := tr.Operation()
		// 断言成HTTP的Transport可以拿到特殊信息

		if ht, ok := tr.(*http.Transport); ok {
			fmt.Println("middleware: ", ht.Request().URL.Query())
		}
		// if ht, ok := tr.(*http.Tranport); ok {
		// 	fmt.Println(ht.Request())
		// }
	}

	return &pb.CreateOrdersReply{
		DtmResult: "SUCCESS",
	}, nil
}

// tcc-cancel
func (s *GoodsService) CreateOrdersCancel(ctx context.Context, req *pb.CreateOrdersRequest) (*pb.CreateOrdersReply, error) {
	s.log.Info("ttc-cancel")
	return &pb.CreateOrdersReply{
		DtmResult: "SUCCESS",
	}, nil
}
