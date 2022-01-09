package impl

import crontab "FrameWork/cronjob"

type worker struct {
	Frequency string
	WorkID    uint32
	State     crontab.TaskState
	behave    func()
}

func (w *worker) UpdateState(state crontab.TaskState) {
	w.State = state
}

func (w *worker) GetSpec() string {
	return w.Frequency
}

func (w *worker) GetWorkId() uint32 {
	return w.WorkID
}

func (w *worker) GetState() crontab.TaskState {
	return w.State
}

func NewWorker(frequency string, workID uint32, behave func()) *worker {
	return &worker{WorkID: workID, Frequency: frequency, behave: behave}
}

func (w *worker) Run() {
	w.behave()
}
