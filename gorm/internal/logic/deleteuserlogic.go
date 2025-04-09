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

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(id int) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	// 初始化响应对象
	resp = &types.CommonResponse{
		Status: http.StatusOK,
		Msg:    "success",
		Data:   nil,
	}
	// 查询用户是否存在
	var user model.User
	if err := l.svcCtx.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户不存在
			resp.Status = http.StatusNotFound
			resp.Msg = "查询记录不存在"
			return resp, nil
		}
		// 数据库查询错误
		resp.Status = http.StatusInternalServerError
		resp.Msg = "Database error: " + err.Error()
		return resp, nil
	}

	// 删除用户记录
	if err := l.svcCtx.DB.Delete(&model.User{}, id).Error; err != nil {
		// 删除失败
		resp.Status = http.StatusInternalServerError
		resp.Msg = "Failed to delete user: " + err.Error()
		return resp, nil
	}
	// 返回成功响应
	resp.Msg = "操作成功"
	return resp, nil

}
