package reliable_event

import (
	"context"
	"reflect"
)

type ReliableEventFactory interface {
	// 注册可靠事件
	RegisterEvent(ctx context.Context, handlerName string, payload reflect.Type, behave func(ctx context.Context), isAsynchronous bool)

	// 运行可靠事件-》异步 or 同步根据参数决定
	StartService(ctx context.Context)

	// 获取某个运行的事件状态
	GetEventByEventId(ctx context.Context)
}

var reliableEventFactory ReliableEventFactory

func InjectReliableEventFactory(impl ReliableEventFactory) {
	reliableEventFactory = impl
}

func RefReliableEventFactory() ReliableEventFactory {
	return reliableEventFactory
}
