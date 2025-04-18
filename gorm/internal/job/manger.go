package job

import (
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobManager 管理多个 Job
type JobManager struct {
	jobs []Job
}

// NewJobManager 创建一个新的 JobManager
func NewJobManager() *JobManager {
	return &JobManager{
		jobs: []Job{},
	}
}

// AddJob 添加 Job 到管理器中
func (m *JobManager) AddJob(job Job) {
	m.jobs = append(m.jobs, job)
}

// Start 启动所有 Job
func (m *JobManager) Start() {
	var wg sync.WaitGroup
	for _, job := range m.jobs {
		wg.Add(1)
		go func(j Job) {
			defer wg.Done()
			logx.Infof("Starting job: %s", j.Name()) // 使用 logx 输出日志
			if err := j.Run(); err != nil {
				logx.Infof("Job %s failed: %v", j.Name(), err)
			} else {
				logx.Infof("Job %s completed successfully", j.Name())
			}
		}(job)
	}
	wg.Wait()
}
