package timer

type Timer interface {
	StartTask(taskName string)
	StopTask(taskName string)
	Remove(taskName string, id int)
	Clear(taskName string)
	Close()
}

type timer struct {
}

func (t *timer) StartTask(taskName string) {
	panic("implement me")
}

func (t *timer) StopTask(taskName string) {
	panic("implement me")
}

func (t *timer) Remove(taskName string, id int) {
	panic("implement me")
}

func (t *timer) Clear(taskName string) {
	panic("implement me")
}

func (t *timer) Close() {
	panic("implement me")
}

func NewTimerTask() Timer {
	return &timer{}
}
