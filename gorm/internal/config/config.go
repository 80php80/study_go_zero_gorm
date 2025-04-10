package config

import "github.com/zeromicro/go-zero/rest"

// DataSource 定义数据源配置
type DataSource struct {
	Driver string `json:"Driver"`
	DSN    string `json:"DSN"`
}

// Redis 配置
type Redis struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	DB       int    `json:"DB"`
}

type Config struct {
	rest.RestConf
	DataSource DataSource `json:"DataSource"`
	Redis      Redis      `json:"Redis"`
}
