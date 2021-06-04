package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}



func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	r.log.Info("data/user: ", po, err)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}