package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/helloMJW/seckill/app/order/service/internal/biz"

	pb "github.com/helloMJW/seckill/api/user/service/v1"
)

var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	data *Data
	log *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (o orderRepo) CreateOrder(ctx context.Context, order *biz.Order) error {

	res, err := o.data.db.Order.Create().SetGid(order.Gid).SetSn("2222").SetUID(333).Save(ctx)

	if err != nil {
		return err
	}

	ucc := pb.NewUserClient(o.data.userRpc)

	res1, err := ucc.GetUser(ctx, &pb.GetUserRequest{Id: 1})

	o.log.Infof("nacos-client-api: ", res1, err)

	o.log.Infof("order-create-result: %v", res)
	return nil
}