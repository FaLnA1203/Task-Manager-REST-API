package task

import "sync"

type TasksList struct {
	tasksList map[string]Task
	mtx       sync.RWMutex
}

func NewTasksList() TasksList {
	return TasksList{
		tasksList: make(map[string]Task),
	}
}

func (l *TasksList) AddTask(task *Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.tasksList[task.Name]
	if ok {
		return TaskAlreadyExists
	}
	l.tasksList[task.Name] = *task
	return nil
}

func (l *TasksList) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasksList))

	for k, v := range l.tasksList {
		tmp[k] = v
	}

	return tmp
}

func (l *TasksList) GetTask(name string) (*Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasksList[name]
	if !ok {
		return nil, TaskNotFound
	}
	return &task, nil
}

func (l *TasksList) ListCompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasksList))

	for k, v := range l.tasksList {
		if v.Completed {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *TasksList) ListUncompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasksList))

	for k, v := range l.tasksList {
		if !v.Completed {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *TasksList) CompliteTask(name string, complete bool) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasksList[name]
	if !ok {
		return TaskNotFound
	}
	if complete {
		task.Complete()
	} else {
		task.Uncomplete()
	}
	l.tasksList[name] = task
	return nil
}

func (l *TasksList) DeleteTask(name string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.tasksList[name]
	if !ok {
		return TaskNotFound
	}
	delete(l.tasksList, name)
	return nil
}
