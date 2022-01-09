package work_pool

import (
	"context"
)

type WorkerFactory interface {
	GetCurrentConsumedRoutineNum(ctx context.Context) int8
	DoBehave(ctx context.Context, behave func(ctx context.Context))
}

var workerFactory WorkerFactory

func RefWorkerFactory() WorkerFactory {
	return workerFactory
}

func InjectWorkerFactory(impl WorkerFactory) {
	workerFactory = impl
}

func init() {
	//todo 获取全局配置
	//初始化当前的配置
	//注入
}
