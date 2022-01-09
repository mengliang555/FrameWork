package impl

import (
	"FrameWork/reliable_event"
	"FrameWork/reliable_event/common"
	"context"
	"reflect"
)

type ReliableEventFactoryImpl struct {
	handlerMap map[string]func(ctx context.Context)
	eventMap   map[string]*common.BaseReliableEventStruct
}

func (r *ReliableEventFactoryImpl) RegisterEvent(ctx context.Context, handlerName string, payload reflect.Type, behave func(ctx context.Context), isAsynchronous bool) {
	//todo remind anything in db
	panic("implement me")
}

func (r *ReliableEventFactoryImpl) StartService(ctx context.Context) {
	//todo Traverse the db with condition:MAX TIME  >= CURRENT TIME AND  CURREN TIME > NEXT RUN TIME
	panic("implement me")
}

func (r *ReliableEventFactoryImpl) GetEventByEventId(ctx context.Context) {
	//TODO SELECT FROM DB WITH EVENT_ID
	panic("implement me")
}

func init() {
	reliable_event.InjectReliableEventFactory(&ReliableEventFactoryImpl{})
}
