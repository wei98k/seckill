package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/peter-wow/seckill/app/goods/service/internal/biz"
	pb "github.com/peter-wow/seckill/api/goods/service/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGoodsService)

type GoodsService struct {
	pb.UnimplementedGoodsServer

	uc *biz.GoodsUsecase
	log *log.Helper
}

func NewGoodsService(uc *biz.GoodsUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		uc: uc,
		log: log.NewHelper(logger),
	}
}

