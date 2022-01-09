package crontab

type TaskState uint8

const (
	TaskReady    = TaskState(0)
	TaskRunning  = TaskState(1)
	TaskStopped  = TaskState(2)
	TaskNotExist = TaskState(10)
)
