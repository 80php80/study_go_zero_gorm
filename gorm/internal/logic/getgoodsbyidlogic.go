package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"

	"gorm/internal/model"
	"gorm/internal/svc"
	"gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdLogic {
	return &GetGoodsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsByIdLogic) GetGoodsById(id int) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	// 初始化返回值
	resp = &types.CommonResponse{
		Status: http.StatusOK,
		Msg:    "Success",
		Data:   nil,
	}
	// 查询用户记录
	var goood model.Good
	if err = l.svcCtx.DB.First(&goood, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Status = http.StatusNotFound
			resp.Msg = "User not found"
			return resp, nil
		}
		resp.Status = http.StatusInternalServerError
		resp.Msg = "Database error: " + err.Error()
		return resp, nil
	}

	resp.Data = goood
	return resp, nil
}
