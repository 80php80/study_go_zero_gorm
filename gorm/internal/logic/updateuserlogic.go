package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm/internal/model"
	"gorm/internal/svc"
	"gorm/internal/types"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.CommonResponse, err error) {
	// 查询用户是否存在
	var user model.User
	result := l.svcCtx.DB.First(&user, req.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 记录不存在，返回友好的提示信息
			return &types.CommonResponse{
				Status: 404,
				Msg:    "记录不存在，请检查 ID 是否正确",
				Data:   nil,
			}, nil
		}
		// 其他数据库错误
		return nil, result.Error
	}

	// 构造更新内容
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Age != nil {
		updates["age"] = *req.Age
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Password != nil {
		updates["password"] = *req.Password
	}

	// 检查是否有需要更新的字段
	if len(updates) == 0 {
		return &types.CommonResponse{
			Status: 400,
			Msg:    "没有需要更新的字段",
			Data:   nil,
		}, nil
	}

	// 执行更新
	result = l.svcCtx.DB.Model(&user).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	// 查询更新后的记录
	err = l.svcCtx.DB.First(&user, req.ID).Error
	if err != nil {
		// 理论上不会发生，但如果查询失败，返回错误
		return nil, err
	}
	// 返回成功响应
	return &types.CommonResponse{
		Status: 200,
		Msg:    "Success",
		Data:   user,
	}, nil
}
