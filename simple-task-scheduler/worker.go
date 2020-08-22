package taskscheduler

import "github.com/google/uuid"

// Worker executes tasks
type Worker struct {
	id                 string
	maxWorkUnits       int
	availableWorkUnits int
	tasks              []Task
}

// NewWorker runs an empty Worker with a max task capacity of zero
func NewWorker(maxWorkUnits int) *Worker {
	return &Worker{
		id:                 uuid.New().String(),
		maxWorkUnits:       maxWorkUnits,
		availableWorkUnits: maxWorkUnits,
		tasks:              []Task{},
	}
}
