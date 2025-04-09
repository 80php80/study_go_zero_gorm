package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm/internal/model"
	"net/http"

	"gorm/internal/svc"
	"gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(id int) (resp *types.CommonResponse, err error) {
	// 初始化返回值
	resp = &types.CommonResponse{
		Status: http.StatusOK,
		Msg:    "Success",
		Data:   nil,
	}
	// 查询用户记录
	var user model.User
	if err := l.svcCtx.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Status = http.StatusNotFound
			resp.Msg = "User not found"
			return resp, nil
		}
		resp.Status = http.StatusInternalServerError
		resp.Msg = "Database error: " + err.Error()
		return resp, nil
	}
	// 返回成功响应
	resp.Data = user
	return resp, nil
}
