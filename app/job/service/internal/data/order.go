package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/job/service/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func (o orderRepo) CreateOrder(ctx context.Context) error {

	return nil
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderQueueRepo {
	return &orderRepo {
		data: data,
		log: log.NewHelper(logger),
	}
}