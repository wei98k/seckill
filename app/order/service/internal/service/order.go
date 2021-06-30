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
	//test: Nacos-client

	//sc := []constant.ServerConfig{
	//	*constant.NewServerConfig("192.168.2.174", 8848),
	//}
	//
	//cc := &constant.ClientConfig{
	//	NamespaceId:         "public", //namespace id
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "/tmp/nacos/log",
	//	CacheDir:            "/tmp/nacos/cache",
	//	RotateTime:          "1h",
	//	MaxAge:              3,
	//	LogLevel:            "debug",
	//}
	//
	//// a more graceful way to create naming client
	//client, err := clients.NewNamingClient(
	//	vo.NacosClientParam{
	//		ClientConfig:  cc,
	//		ServerConfigs: sc,
	//	},
	//)
	//
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//r := registry.New(client)
	//
	//
	//
	//conn, err := grpc.DialInsecure(
	//	context.Background(),
	//	grpc.WithEndpoint("discovery:///user.grpc"),
	//	grpc.WithDiscovery(r),
	//)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("nacos-client-conn: ", conn, err)
	//
	//ucc := upb.NewUserClient(conn)
	//
	//res, err := ucc.GetUser(ctx, &upb.GetUserRequest{
	//	Id: 1,
	//})
	//
	//log.Println("nacos-client-api: ", res, err)

	return &pb.CreateOrderReply{},nil
}