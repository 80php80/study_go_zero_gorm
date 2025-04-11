package job

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm/internal/model"
	"gorm/internal/svc"
	"math/rand"
	"sync"
	"time"
)

type BulkInsertJob struct {
	svcCtx *svc.ServiceContext
}

func NewBulkInsertJob(svcCtx *svc.ServiceContext) *BulkInsertJob {
	return &BulkInsertJob{
		svcCtx: svcCtx,
	}
}

func (c *BulkInsertJob) Name() string {
	return "BulkInsertJob"
}

func (j *BulkInsertJob) Run() error {
	batchSize := 1000         // 每次插入的数据
	totalRecords := 1_000_000 //总记录数
	newWorkers := 10          // 并发协程数

	taskChan := make(chan []model.Good, newWorkers)

	var wg sync.WaitGroup //WaitGroup 用于等待所有 Goroutine 完成

	// 启动多个 Goroutine 处理插入任务
	for i := 0; i < newWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range taskChan {
				if err := j.svcCtx.DB.Create(&batch).Error; err != nil {
					logx.Errorf("Failed to insert batch: %v", err)
					continue
				}
			}
		}()
	}
	//记录开始时间
	startTime := time.Now()
	for i := 0; i < totalRecords; i += batchSize {
		var goods []model.Good
		for j := 0; j < batchSize && i+j < totalRecords; j++ {
			goods = append(goods, model.Good{
				//Name:  fmt.Sprintf("good_%d", i),
				Name:  fmt.Sprintf("商品_%d", i+j),
				Price: float64(rand.Intn(10000) + 1), // 随机生成价格
			})
		}
		taskChan <- goods
	}
	//关闭channel 通知go 协程 任务完成
	close(taskChan)
	wg.Wait() // 等待所有的 Goroutine  完成
	// 计算总耗时
	elapsed := time.Since(startTime)
	logx.Infof("Finished inserting %d records in %v", totalRecords, elapsed)

	return nil

}
