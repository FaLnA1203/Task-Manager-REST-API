package task

import "errors"

var TaskNotFound = errors.New("task not found")
var TaskAlreadyExists = errors.New("task already exists")
