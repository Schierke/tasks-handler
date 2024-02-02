package handler

import "errors"

var (
	Timeout             = "request time out"
	ErrTaskIdNotFound   = "can't get the task id"
	ErrUpdateOpsMember  = "can't update ops memeber"
	ErrLocationNotFound = "can't found the task location"
)

var (
	ErrTasksNotFound       = errors.New("can't get list of tasks")
	ErrFilterValueNotFound = errors.New("can't get the filter value")
)
