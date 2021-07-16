package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Gid int64
	Amount int64
	Sn string
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *Order) error
}

type OrderUsecase struct {
	repo OrderRepo
	log *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderUsecase) Create(ctx context.Context, order *Order) error {
	return uc.repo.CreateOrder(ctx, order)
}

type OrderQueueRepo interface {
	CreateOrder(ctx context.Context) error
}

type OrderQueueUsecase struct {
	repo OrderQueueRepo
	log *log.Helper
}

func NewOrderQueueUsecase(repo OrderQueueRepo, logger log.Logger) *OrderQueueUsecase {
	return &OrderQueueUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderQueueUsecase) CreateQueue(ctx context.Context) error {
	uc.log.Info("queue test")
	return uc.repo.CreateOrder(ctx)
}
