package common

import "time"

type EventStatus uint8

const (
	Init    = EventStatus(1)
	Running = EventStatus(2)
	Pause   = EventStatus(3)
	Stop    = EventStatus(4)
	Remove  = EventStatus(5)
)

type EventStruct struct {
	Id          int
	EventId     int
	EventStatus EventStatus
	HadTryTime  int
	MaxTryTime  int
	CreateTime  time.Time
	ModifyTime  time.Time
}
