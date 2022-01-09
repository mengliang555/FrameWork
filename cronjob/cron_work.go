package crontab

type CronWork interface {
	Run()
	GetWorkId() uint32
	GetState() TaskState
	GetSpec() string
	UpdateState(TaskState)
}
