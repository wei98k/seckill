
package service

import (
	"context"
	pb "github.com/helloMJW/seckill/api/user/service/v1"
	"github.com/helloMJW/seckill/app/user/service/internal/biz"
)


func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	u := &biz.RegUser{Username: req.Username, PasswordHash: req.Password}
	rv, err := s.uc.Create(ctx, u)

	s.log.Infof("service/user/CreateUser: rv: %v, err: %v", rv, err)

	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	u := &biz.UpdateUser{
		PasswordHash: req.Password,
	}

	err := s.uc.Update(ctx, req.Id, u)

	if err != nil {
		// 返回错误信息
		return nil, err
	}
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	s.log.Info("getuser request: ", rv, req)
	return &pb.GetUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, err
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
