package job

type Job interface {
	Run() error   // 执行的方法
	Name() string //job 名称
}
