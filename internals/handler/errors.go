package handler

import "errors"

var (
	Timeout           = "request time out"
	ErrTaskIdNotFound = "can't get the task id"
)

var (
	ErrTasksNotFound = errors.New("can't get list of tasks")

	ErrFilterValueNotFound = errors.New("can't get the filter value")
)
