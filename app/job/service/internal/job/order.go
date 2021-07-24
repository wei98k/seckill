package job

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/job/service/internal/conf"
	"github.com/peter-wow/seckill/app/job/service/internal/service"
)

type OrderJob struct {

}

// NewOrderJobServer new a job server.
func NewOrderJobServer(c *conf.Server, order *service.OrderService, logger log.Logger) *OrderJob {

	fmt.Println(c.Http.Addr)

	logger.Log(log.LevelInfo, c.Http.Addr)

	order.Start()

	return &OrderJob{}
}
