package task

type TasksList struct {
	TasksList map[string]Task
}

func NewTasksList() TasksList {
	return TasksList{
		TasksList: make(map[string]Task),
	}
}

func (tl *TasksList) AddTask(task *Task) error {
	_, ok := tl.TasksList[task.Name]
	if ok {
		return TaskAlreadyExists
	}
	tl.TasksList[task.Name] = *task
	return nil
}

func (l *TasksList) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.TasksList))

	for k, v := range l.TasksList {
		tmp[k] = v
	}

	return tmp
}

func (tl *TasksList) CompliteTask(name string) error {
	task, ok := tl.TasksList[name]
	if !ok {
		return TaskNotFound
	}
	task.Complete()
	tl.TasksList[name] = task
	return nil
}

func (tl *TasksList) DeleteTask(name string) error {
	_, ok := tl.TasksList[name]
	if !ok {
		return TaskNotFound
	}
	delete(tl.TasksList, name)
	return nil
}
