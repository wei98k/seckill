package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/helloMJW/seckill/app/user/service/internal/biz"
	"github.com/helloMJW/seckill/pkg/utils/encryption"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}
//
// TODO::用户名不可以重复 手机号不可重复 邮箱不可以重复
// 创建用户
//
func (r *userRepo) CreateUser(ctx context.Context, user *biz.RegUser) (id int64, err error) {

	//hp, err := encryption.Md5Password(user.PasswordHash)
	//r.log.Info("密码: ", hp, err)
	//if err != nil {
	//	return 0, err
	//}

	po, err := r.data.db.User.Create().SetUsername(user.Username).SetPasswordHash(encryption.Md5Password(user.PasswordHash)).Save(ctx)

	r.log.Infof("data/user/CreateUser: po: %v, err: %v", po, err)

	return po.ID, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)

	// test start
	// r.data.rdb.Set(ctx, "hello", "kratos", 0).Result()
	//str, err1 := r.data.rdb.Get(ctx, "hello").Result()
	//fmt.Println("redis-test-get :", str, err1)
	// test stop
	r.log.Info("data/user: ", po, err)

	// test kafka start
	//s := NewPublisher("test", []string{"127.0.0.1:9092"})
	// test kafka stop

	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

