package data

import (
	"context"
	"database/sql"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/peter-wow/seckill/app/order/service/internal/conf"
	"github.com/peter-wow/seckill/app/order/service/internal/data/ent"
	"github.com/peter-wow/seckill/app/order/service/internal/data/kafka"
	"github.com/peter-wow/seckill/app/order/service/internal/data/kafka/event"
	ggrpc "google.golang.org/grpc"
	"strings"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo, NewSeckillOrderRepo, NewSeckillGoodsRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *ent.Client

	msql *sql.DB

	// Gconn *grpc.ClientConn
	userRpc *ggrpc.ClientConn

	rdb *redis.Client

	kafka event.Sender
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger, rr *registry.Registry) (*Data, func(), error) {

	//orm ent
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

	// mysql
	msql, err := sql.Open("mysql", conf.Database.Source)

	if err != nil {
		log.Errorf("failed conn mysql %v", err)
	}

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	//gRpc
	userRpc, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.grpc"),
		grpc.WithDiscovery(rr),
	)

	if err != nil {
		panic("grpc-error")
	}

	//kafka-sender
	senderClient, err1 := kafka.NewKafkaSender(strings.Split(conf.Kafka.Addr, ","), "order")

	if err1 != nil {
		panic("kafka-error")
	}

	d := &Data{
		db: client,
		msql: msql,
		userRpc: userRpc,
		rdb: rdb,
		kafka: senderClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
