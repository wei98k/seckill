package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/job/service/internal/conf"
	"github.com/peter-wow/seckill/app/job/service/internal/service"
	"github.com/peter-wow/seckill/app/job/service/job"
)

func NewJOBServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *job.Server  {
	s := job.NewKafkaReceiver([]string{"192.168.2.27:9092"}, "order")

	return s
}