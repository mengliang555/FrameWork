package common

import (
	"context"
	"reflect"
)

type ReliableEvent interface {
	GetCurrentStatus()
	StopEvent()
	IsAsynchronous() bool
}

type BaseReliableEventStruct struct {
	Payload        reflect.Type
	HandlerName    string
	EventBehave    func(ctx context.Context)
	IsAsynchronous bool
}
