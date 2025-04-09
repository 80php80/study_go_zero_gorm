package logic

import (
	"context"

	"gorm/internal/svc"
	"gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GormLogic {
	return &GormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GormLogic) Gorm(req *types.GoodsBase) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
