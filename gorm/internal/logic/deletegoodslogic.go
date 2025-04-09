package logic

import (
	"context"

	"gorm/internal/svc"
	"gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsLogic {
	return &DeleteGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGoodsLogic) DeleteGoods() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
