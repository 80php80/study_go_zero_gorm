package config

import "github.com/zeromicro/go-zero/rest"

// DataSource 定义数据源配置
type DataSource struct {
	Driver string `json:"Driver"`
	DSN    string `json:"DSN"`
}

type Config struct {
	rest.RestConf
	DataSource DataSource `json:"DataSource"`
}
