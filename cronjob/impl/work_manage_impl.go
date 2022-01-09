package impl

import (
	crontab "FrameWork/cronjob"
	"github.com/robfig/cron"
	"sync"
)

type workRunManage struct {
	lock     *sync.RWMutex
	work     crontab.Work
	workCron *cron.Cron
}

type workRunFactory struct {
	size     int32
	lock     *sync.RWMutex
	workList map[uint32]*workRunManage
}

func (w *workRunFactory) ContinueTargetWork(id uint32) {
	if !w.ContainWork(id) {
		panic("not contain the job")
	}
	w.workList[id].workCron.Start()
	w.workList[id].work.UpdateState(crontab.TaskRunning)
}

func (w *workRunFactory) ContainWork(id uint32) bool {
	w.lock.RLock()
	defer w.lock.RUnlock()
	if _, ok := w.workList[id]; ok {
		return true
	}
	return false
}

func (w *workRunFactory) InitCronWorkFactoryCapacity(capacity int32) crontab.CronWorkFactory {
	w.size = capacity
	w.lock = &sync.RWMutex{}
	w.workList = make(map[uint32]*workRunManage)
	return w
}

func (w *workRunFactory) RegisterWork(work crontab.Work) crontab.CronWorkFactory {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.size == 0 {
		panic("the factory should be initialize")
	}
	work.UpdateState(crontab.TaskReady)
	if w.workList[work.GetWorkId()] != nil {
		w.workList[work.GetWorkId()].workCron.Stop()
	}
	w.workList[work.GetWorkId()] = &workRunManage{workCron: cron.New(), work: work, lock: &sync.RWMutex{}}
	return w
}

func (w *workRunFactory) GetWorkStateByIdentity(id uint32) crontab.TaskState {
	w.workList[id].lock.RLock()
	defer w.workList[id].lock.RUnlock()
	if v, ok := w.workList[id]; ok {
		return v.work.GetState()
	} else {
		return crontab.TaskNotExist
	}
}

func (w *workRunFactory) StartWorkByByIdentity(id uint32) crontab.CronWorkFactory {
	if !w.ContainWork(id) {
		panic("not contain job")
	}
	w.workList[id].lock.Lock()
	defer w.workList[id].lock.Unlock()
	err := w.workList[id].workCron.AddJob(w.workList[id].work.GetSpec(), w.workList[id].work)
	if err != nil {
		panic(err)
	}
	w.workList[id].workCron.Start()
	w.workList[id].work.UpdateState(crontab.TaskRunning)
	return w
}

func (w *workRunFactory) StopWorkByByIdentity(id uint32) crontab.CronWorkFactory {
	if !w.ContainWork(id) {
		panic("not contain job")
	}

	w.workList[id].lock.Lock()
	defer w.workList[id].lock.Unlock()
	w.workList[id].workCron.Stop()
	w.workList[id].work.UpdateState(crontab.TaskStopped)
	return w
}

func (w *workRunFactory) RevokeWorkByByIdentity(id uint32) crontab.CronWorkFactory {
	w.workList[id].lock.Lock()
	defer w.workList[id].lock.Unlock()
	if w.GetWorkStateByIdentity(id) != crontab.TaskStopped {
		w.StopWorkByByIdentity(id)
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	delete(w.workList, id)
	return w
}

func init() {
	crontab.InjectCronWorkFactory(&workRunFactory{})
}
