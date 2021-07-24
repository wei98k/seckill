package job

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/peter-wow/seckill/app/job/service/internal/service"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	o service.OrderService
}

func (s Server) Start(ctx context.Context) error {
	s.o.Start()
	return nil
}

func (s Server) Stop(ctx context.Context) error {

	return nil
}

