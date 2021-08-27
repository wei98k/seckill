package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/goods/service/internal/biz"
)

var _ biz.OrdersRepo = (*ordersRepo)(nil)

type ordersRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrdersRepo(data *Data, logger log.Logger) biz.OrdersRepo {
	return &ordersRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "moudle", "data/orders")),
	}
}

func (m ordersRepo) GetOrders(ctx context.Context, id int64) (*biz.Orders, error) {
	result, err := m.data.db.Order.Get(ctx, id)

	if err != nil {
		m.log.Error(err)
		return nil, err
	}

	return &biz.Orders{
		Sn: result.Sn,
	}, nil
}

func (m ordersRepo) ListOrders(ctx context.Context) (*biz.Orders, error) {
	return &biz.Orders{}, nil
}
func (m ordersRepo) CreateOrders(ctx context.Context, orders biz.Orders) error {

	_, err := m.data.db.Order.Create().
		SetSn("ddd").
		SetGid(222).
		SetUID(99).
		Save(ctx)

	if err != nil {
		m.log.Info(err)
		return err
	}
	return nil
}
func (m ordersRepo) UpdateOrders(ctx context.Context, orders biz.Orders) error {
	return nil
}
