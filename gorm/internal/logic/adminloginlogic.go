package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm/internal/model"
	"gorm/internal/utils"
	"strings"
	"time"

	"gorm/internal/svc"
	"gorm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminLoginRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	var user model.User
	result := l.svcCtx.DB.Where(&model.User{Email: req.Email}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &types.CommonResponse{
				Status: 404,
				Msg:    "用户不存在或者邮箱不正确",
				Data:   nil,
			}, nil
		}
	}
	if strings.TrimSpace(user.Password) != strings.TrimSpace(req.Password) {
		return &types.CommonResponse{
			Status: 400,
			Msg:    "密码不正确",
			Data:   nil,
		}, nil
	}

	// 生成jwt
	token, err := utils.CreatedJwtToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	// 把 token存到redis
	userIDStr := fmt.Sprintf("%d", user.ID)
	expiration := time.Hour * 1 // token 过期时间
	err = l.svcCtx.RedisClient.Set(context.Background(), userIDStr, token, expiration).Err()
	if err != nil {
		return &types.CommonResponse{
			Status: 500,
			Msg:    "Failed to store user session in Redis",
			Data:   nil,
		}, nil
	}
	// 构造返回数据
	loginResponse := types.AdminLoginResponse{
		Token: token,
		User: types.UserBase{
			Name:  user.Name,
			Age:   user.Age,
			Email: user.Email,
		},
	}
	// 返回成功响应
	return &types.CommonResponse{
		Status: 200,
		Msg:    "登录成功",
		Data:   loginResponse,
	}, nil
}
