package logic

import (
	"context"
	"gorm/internal/model"

	"gorm/internal/svc"
	"gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.AddUserRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	user := &model.User{
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	// 插入数据到数据库
	result := l.svcCtx.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 返回成功响应
	return &types.CommonResponse{
		Status: 200,
		Msg:    "Success",
		Data:   user,
	}, nil

}
