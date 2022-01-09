package crontab

type Work interface {
	Run()
	GetWorkId() uint32
	GetState() TaskState
	GetSpec() string
	UpdateState(TaskState)
}
