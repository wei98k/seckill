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

func (r *userRepo) UpdateUser(ctx context.Context, id int64, user *biz.UpdateUser) error {
	ro, err := r.data.db.User.UpdateOneID(id).
			SetPasswordHash(encryption.Md5Password(user.PasswordHash)).Save(ctx)
	r.log.Infof("update-user-result: %v", ro)
	return err
}

func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	return r.data.db.User.DeleteOneID(id).Exec(ctx)
}
//
// TODO::用户名不可以重复 手机号不可重复 邮箱不可以重复
// 创建用户
//
func (r *userRepo) CreateUser(ctx context.Context, user *biz.RegUser) (id int64, err error) {

	po, err := r.data.db.User.Create().SetUsername(user.Username).SetPasswordHash(encryption.Md5Password(user.PasswordHash)).Save(ctx)

	r.log.Infof("data/user/CreateUser: po: %v, err: %v", po, err)

	return po.ID, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	r.log.Info("data/user: ", po, err)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}
