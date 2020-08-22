package taskscheduler

import (
	"fmt"

	"github.com/google/uuid"
)

// Worker executes tasks
type Worker struct {
	id                 string
	maxWorkUnits       int
	availableWorkUnits int
	tasks              []*Task
}

// NewWorker runs an empty Worker with a max task capacity of zero
func NewWorker(maxWorkUnits int) *Worker {
	return &Worker{
		id:                 uuid.New().String(),
		maxWorkUnits:       maxWorkUnits,
		availableWorkUnits: maxWorkUnits,
		tasks:              []*Task{},
	}
}

// AvailableWorkUnits returns the number of available work units that can be
// used for tasks
func (w Worker) AvailableWorkUnits() int {
	return w.availableWorkUnits
}

// AddTask will add a new Task to a Worker
//
// An error will be returned if the task fails to be added
func (w *Worker) AddTask(t *Task) error {
	if t.unitSize <= w.maxWorkUnits {
		w.tasks = append(w.tasks, t)
		t.status = SCHEDULED
	}

	return fmt.Errorf("not able to add task")
}
