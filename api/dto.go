package api

import (
	"encoding/json"
	"errors"
	"time"
)

type DtoTask struct {
	Name        string `json:"name"`
	Description string `json: "description"`
}

func (task *DtoTask) Validate() error {
	if task.Name == "" {
		return errors.New("name is required")
	}
	if task.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

type DtoError struct {
	Error string `json: "error"`
	Time  time.Time
}

func (e *DtoError) ErrorToString() string {
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

type DtoCompleted struct {
	Completed bool `json: "completed"`
}
