package service

import (
	"context"

	pb "github.com/peter-wow/seckill/api/order/service/v1"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
)

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {

	s.log.Infof("CreateOrder request: %v", req)

	// s.goods.SetSeckillGoodsOver(ctx, 99)
	// s.goods.GetSeckillGoodsOver(ctx, 98)

	//data := &biz.Order{
	//	Gid: req.Gid,
	//	Amount: 1,
	//}
	//
	//err := s.uc.Create(ctx, data)
	//if err != nil {
	//	return nil, err
	//}

	//##### kafka send qps/9.06

	// post := &biz.SeckillOrder{
	// 	UserId: 99,
	// 	GoodsId: req.Gid,
	// 	OrderId: 77,
	// }

	// err := s.so.SendKafka(ctx, post)

	// if err != nil {
	// 	fmt.Println(err)
	// 	s.log.Debug(err)
	// 	// s.log.Error(err)
	// }

	//===== 测试mysql get qps/1006.57
	//data, err := s.so.GetOrder(ctx, 2439)
	//
	//if err != nil {
	//	s.log.Error(err)
	//}
	//
	//s.log.Info(data)

	//====== 测试mysql create  qps/440.46
	//post := &biz.SeckillOrder{
	//	UserId: 99,
	//	GoodsId: req.Gid,
	//	OrderId: 77,
	//}
	//
	//s.so.PostOrder(ctx, post)

	//===== 测试分布式管理器dtm start

	// saga := dtmcli.NewSaga(DtmServer, dtmcli.MustGenGid(DtmServer)).
	// 	Add("", "")

	// err := saga.Submit()
	// dtmcli.FatalIfError(err)

	// return saga.Gid

	//===== 测试分布式管理器dtm end

	return &pb.CreateOrderReply{}, nil
}

func (s *OrderService) CreateSeckillOrder(ctx context.Context, req *pb.CreateSeckillOrderRequest) (*pb.CreateSeckillOrderReply, error) {
	//当用户请求走到这里的时候、说明用户信息已校验完成、合法用户、合法请求
	//这儿的操作需要保持原子性
	s.log.Info("input data %v ", req)
	//TODO::检验参数-交给中间件、验证器处理
	//goods_id cache not found 说明非法ID

	//验证秒杀活动
	//goods, err1 := s.goods.GetSeckillGoods(ctx, req.Gid)
	//
	//if err1 != nil {
	//	s.log.Errorf("get goods error %v", err1)
	//	return nil, err1
	//}
	//s.log.Infof("get goods data: %v", goods)

	//TODO::step1: 验证活动是否开始或结束

	//step2: 验证当前活动商品库存
	if s.goods.GetSeckillGoodsOver(ctx, req.Gid) {
		param := &biz.SeckillOrder{
			UserId:  88,
			GoodsId: req.Gid,
			OrderId: 99,
		}

		err := s.so.CreateSeckillOrder(ctx, param)
		s.log.Info("seckill-order-service: err ", err)
	}
	// s.goods.SetSeckillGoodsOver(ctx, req.Gid)
	//TODO::处理系统反馈信息、转换普通用户可见信息
	return &pb.CreateSeckillOrderReply{}, nil
	// 状态码返回
}
