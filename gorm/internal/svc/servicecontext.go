package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/internal/config"
	"gorm/internal/middleware"
	"net/http"
)

type ServiceContext struct {
	Config            config.Config
	DB                *gorm.DB
	JwtAuthMiddleware func(http.HandlerFunc) http.HandlerFunc // 注册中间件
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 GORM 数据库连接
	db, err := gorm.Open(mysql.Open(c.DataSource.DSN), &gorm.Config{})
	if err != nil {
		logx.Errorf("Failed to connect to database: %v", err)
		panic(err)
	}

	logx.Info("Database connected successfully!")
	// 创建 JwtAuthMiddleware 实例
	jwtMiddleware := middleware.NewJwtAuthMiddleware("your-secret-key")

	return &ServiceContext{
		Config:            c,
		DB:                db,
		JwtAuthMiddleware: jwtMiddleware.Handle,
	}
}
