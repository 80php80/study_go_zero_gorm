package logic

import (
	"context"

	"grpc/goodsrpc"
	"grpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckGoodsLogic {
	return &CheckGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckGoodsLogic) CheckGoods(in *goodsrpc.CheckGoodsRequest) (*goodsrpc.CheckGoodsResponse, error) {
	// todo: add your logic here and delete this line
	if in.GoodsId == 123 {
		return &goodsrpc.CheckGoodsResponse{
			IsAvailable: true,
			Message:     "商品可用",
		}, nil
	}

	return &goodsrpc.CheckGoodsResponse{
		IsAvailable: false,
		Message:     "商品不可用",
	}, nil
}
