package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/nacos/registry"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/helloMJW/seckill/app/order/service/internal/conf"
	"github.com/helloMJW/seckill/app/order/service/internal/data/ent"
	ggrpc "google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *ent.Client

	// Gconn *grpc.ClientConn
	userRpc *ggrpc.ClientConn
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger, rr *registry.Registry) (*Data, func(), error) {

	log := log.NewHelper(log.With(logger, "module", "server-service/data"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	userRpc, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.grpc"),
		grpc.WithDiscovery(rr),
	)

	if err != nil {
		panic("grpc-error")
	}

	d := &Data{
		db: client,
		userRpc: userRpc,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
