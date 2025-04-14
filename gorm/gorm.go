package main

import (
	"flag"
	"fmt"
	"gorm/internal/job"

	"gorm/internal/config"
	"gorm/internal/handler"
	"gorm/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gorm-api.yaml", "the config file")

func main() {

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 初始化 Job 管理器
	jobManager := job.NewJobManager()
	// 添加 Job 到管理器中
	jobManager.AddJob(job.NewBulkInsertJob(ctx))
	jobManager.AddJob(job.NewCheckGoodsJob(ctx))

	// 异步启动所有 Job
	go func() {
		fmt.Println("Starting all jobs...")
		jobManager.Start()
	}()

	fmt.Printf("Stasrting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
