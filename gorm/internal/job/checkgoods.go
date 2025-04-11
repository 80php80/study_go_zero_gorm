package job

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm/internal/svc"
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
	for i := 0; i < 100; i++ {
		logx.Infof("这是第二个任务 %d", i)
	}
	return nil

}
