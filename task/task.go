package task

import (
	"time"
)

type Task struct {
	Name          string
	Description   string
	Completed     bool
	TimeCreated   time.Time
	TimeCompleted *time.Time
}

func CreateTask(name string, description string) Task {
	return Task{
		Name:        name,
		Description: description,
		Completed:   false,

		TimeCreated:   time.Now(),
		TimeCompleted: nil,
	}
}

func (t *Task) Complete() {
	TimeCompleted := time.Now()
	t.Completed = true
	t.TimeCompleted = &TimeCompleted
}
