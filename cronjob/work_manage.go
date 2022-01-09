package crontab

type CronWorkFactory interface {
	InitCronWorkFactoryCapacity(int32) CronWorkFactory
	RegisterWork(work Work) CronWorkFactory
	GetWorkStateByIdentity(id uint32) TaskState
	ContinueTargetWork(id uint32)
	StartWorkByByIdentity(id uint32) CronWorkFactory
	StopWorkByByIdentity(id uint32) CronWorkFactory
	RevokeWorkByByIdentity(id uint32) CronWorkFactory
	ContainWork(id uint32) bool
}

var cronWorkFactory CronWorkFactory

func InjectCronWorkFactory(factory CronWorkFactory) {
	cronWorkFactory = factory
}

func GetCronWorkFactoryImpl() CronWorkFactory {
	return cronWorkFactory
}
