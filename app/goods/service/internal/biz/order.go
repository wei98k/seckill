package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// 定义表结构体
type Orders struct {
	Id  int64
	Sn  string
	Uid int64
}

type OrderGoods struct {
	Id         int64
	OrderId    int64
	GoodsTitle string
}

type OrdersUsecase struct {
	repo OrdersRepo
	log  *log.Helper
}

// 定义注入方法
func NewOrdersUsecase(repo OrdersRepo, logger log.Logger) *OrdersUsecase {
	return &OrdersUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// 定义接口
type OrdersRepo interface {
	GetOrders(ctx context.Context, id int64) (*Orders, error)
	ListOrders(ctx context.Context) (*Orders, error)
	CreateOrders(ctx context.Context, orders Orders) error
	UpdateOrders(ctx context.Context, orders Orders) error
}

// 提供给service调用的方法
func (m *OrdersUsecase) GetOrders(ctx context.Context, id int64) (*Orders, error) {
	return m.repo.GetOrders(ctx, id)
}

func (m *OrdersUsecase) CreateOrders(ctx context.Context, orders Orders) error {
	return m.repo.CreateOrders(ctx, orders)
}
