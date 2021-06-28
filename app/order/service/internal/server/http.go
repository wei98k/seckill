package server

import (
	pb "github.com/helloMJW/seckill/api/order/service/v1"
	"github.com/helloMJW/seckill/app/order/service/internal/conf"
	"github.com/helloMJW/seckill/app/order/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, s *service.OrderService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
	)
	srv.HandlePrefix("/", pb.NewOrderHandler(s, m))
	return srv
}
