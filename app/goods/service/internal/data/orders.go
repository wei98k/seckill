package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	rr "github.com/go-resty/resty/v2"
	"github.com/peter-wow/seckill/app/goods/service/internal/biz"
	"github.com/yedf/dtmcli"
)

var _ biz.OrdersRepo = (*ordersRepo)(nil)

type TransReq struct {
	Sn string `json:"sn"`
}

type TransReq1 struct {
	Gid    int64 `json:"gid"`
	Amount int64 `json:"amount"`
}

type ordersRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrdersRepo(data *Data, logger log.Logger) biz.OrdersRepo {
	return &ordersRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "moudle", "data/orders")),
	}
}

func (m ordersRepo) GetOrders(ctx context.Context, id int64) (*biz.Orders, error) {
	result, err := m.data.db.Order.Get(ctx, id)

	if err != nil {
		m.log.Error(err)
		return nil, err
	}

	return &biz.Orders{
		Sn: result.Sn,
	}, nil
}

func (m ordersRepo) ListOrders(ctx context.Context) (*biz.Orders, error) {
	return &biz.Orders{}, nil
}

// 创建普通商品订单-try阶段
func (m ordersRepo) CreateOrders(ctx context.Context, orders biz.Orders) error {

	// 检查用户合法性
	// 检查优惠卷是否合法
	// m.log.Info("model-goods-create-orders-start")
	// if tr, ok := transport.FromServerContext(ctx); ok {
	// 	// kind := tr.Kind().String()
	// 	// operation := tr.Operation()
	// 	// 断言成HTTP的Transport可以拿到特殊信息
	// 	m.log.Info("model-goods-create-orders-ing:step1")
	// 	if ht, ok := tr.(*http.Transport); ok {
	// 		m.log.Info("model-goods-create-orders-ing:step2")
	// 		fmt.Println("middleware: ", ht.Request().URL.Query())

	// 		barrier, err := dtmcli.BarrierFromQuery(ht.Request().URL.Query())
	// 		dtmcli.FatalIfError(err)

	// 		m.log.Debug("dtm-result: ", dtmcli.ResultSuccess)
	// 		m.log.Debug("dtm-result-error: ", err)

	// 		barrier.Call(BeginTx(m.data.sql), func(db dtmcli.DB) error {
	// 			// return adjustTrading(db, transInUID, req.Amount)
	// 			m.log.Info("model-goods-create-orders-ing:step3")
	// 			// insert into goods.orders (`sn`, `uid`, `gid`, `created_at`, `updated_at`) values ("mmm11111", 99, 66, "2021-08-29 16:12:17", "2021-08-29 16:12:17");
	// 			sn := "mmm11111"
	// 			uid := 99
	// 			gid := 77
	// 			ctime := "2021-08-29 16:12:17"
	// 			utime := "2021-08-29 16:12:17"

	// 			affected, err := dtmcli.DBExec(db, "insert into goods.orders (`sn`, `uid`, `gid`, `created_at`, `updated_at`) values (?, ?, ?, ?, ?)", sn, uid, gid, ctime, utime)

	// 			m.log.Debug("dtmcli-exec: ", affected)
	// 			m.log.Debug("dtmcli-exec-error: ", err)
	// 			return err
	// 		})

	// 	}
	// 	// if ht, ok := tr.(*http.Tranport); ok {
	// 	// 	fmt.Println(ht.Request())
	// 	// }
	// }

	// _, err := m.data.db.Order.Create().
	// 	SetSn("ddd").
	// 	SetGid(222).
	// 	SetUID(99).
	// 	Save(ctx)

	// if err != nil {
	// 	m.log.Info(err)
	// 	return err
	// }

	//分布式事务调用

	var DtmServer string = "http://127.0.0.1:8080/api/dtmsvr"
	var Busi string = "http://127.0.0.1:8003"

	var Busi1 string = "http://127.0.0.1:8001"

	gid := dtmcli.MustGenGid(DtmServer)
	err := dtmcli.TccGlobalTransaction(DtmServer, gid, func(tcc *dtmcli.Tcc) (*rr.Response, error) {
		resp, err := tcc.CallBranch(&TransReq{Sn: "sss123123"}, Busi+"/ordersTry", Busi+"/ordersConfirm", Busi+"/ordersCancel")
		if err != nil {
			return resp, err
		}

		return tcc.CallBranch(&TransReq1{Gid: 111, Amount: 99}, Busi1+"/seckill/orderTry", Busi1+"/seckill/orderConfirm", Busi1+"/seckill/orderCancel")
	})
	m.log.Debug("dtm: ", err)
	// dtmcli.FatalIfError(err)

	return nil
}

// 创建普通商品订单-Confirm阶段
func (m ordersRepo) CreateOrdersConfirm() (interface{}, error) {

	// 直接生成订单、插入数据库

	res := dtmcli.OrString("ok", "happy", "SUCCESS")
	// dtmcli.Logf("%s %s result: %s", busi, info.String(), res)
	type M = map[string]interface{}

	return M{"dtm_result": res}, nil
}

// 创建普通商品订单-Cancel阶段

func (m ordersRepo) UpdateOrders(ctx context.Context, orders biz.Orders) error {
	// 插入失败, 取消操作
	return nil
}
