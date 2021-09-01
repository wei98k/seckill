package data

import (
	"context"
	"database/sql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/peter-wow/seckill/app/goods/service/internal/conf"
	"github.com/peter-wow/seckill/app/goods/service/internal/data/ent"
	"github.com/yedf/dtmcli"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGoodsRepo, NewOrdersRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *ent.Client

	sql *sql.DB
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
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

	//sql-dtm
	// 结构体转为 map[string][string]
	var dbSqlDtmConf = map[string]string{
		"driver":   conf.Database.Driver,
		"host":     "192.168.0.111",
		"user":     "root",
		"password": "Root@123456",
		"port":     "3307",
	}
	dbSqlDtm, err := dtmcli.SdbGet(dbSqlDtmConf)
	dtmcli.FatalIfError(err)

	d := &Data{
		db:  client,
		sql: dbSqlDtm,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

// 开启mysql事务
func BeginTx(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	dtmcli.FatalIfError(err)
	return tx
}
