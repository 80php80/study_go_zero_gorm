package job

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm/internal/svc"
	"sync"
)

type CheckGoodsJob struct {
	svcCtx *svc.ServiceContext
}

func NewCheckGoodsJob(svcCtx *svc.ServiceContext) *CheckGoodsJob {
	return &CheckGoodsJob{
		svcCtx: svcCtx,
	}
}

func (j *CheckGoodsJob) Name() string {
	return "CheckGoodsJob"
}

func (j *CheckGoodsJob) Run() error {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			logx.Infof("这是第二个任务 %d", num)
		}(i)
	}
	wg.Wait()

	return nil

}
