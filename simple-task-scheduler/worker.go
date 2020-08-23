package taskscheduler

import (
	"fmt"

	"github.com/hashicorp/go-hclog"

	"github.com/google/uuid"
)

type WorkerOptions struct {
	ID     string
	Logger hclog.Logger
}

// Worker executes tasks
type Worker struct {
	id                 string
	maxWorkUnits       int
	availableWorkUnits int
	tasks              []*Task
	logger             hclog.Logger
}

// NewWorker runs an empty Worker with a max task capacity of zero
func NewWorker(maxWorkUnits int, options *WorkerOptions) *Worker {
	if options == nil {
		options = &WorkerOptions{
			ID:     uuid.New().String(),
			Logger: hclog.NewNullLogger(),
		}
	}

	if options.ID == "" {
		options.ID = uuid.New().String()
	}

	return &Worker{
		id:                 options.ID,
		maxWorkUnits:       maxWorkUnits,
		availableWorkUnits: maxWorkUnits,
		tasks:              []*Task{},
		logger:             options.Logger,
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
	if t.unitSize <= w.availableWorkUnits {
		w.tasks = append(w.tasks, t)
		t.status = SCHEDULED
		w.availableWorkUnits -= t.unitSize
		w.logger.Info("task added", "task_id", t.id, "available_work_units", w.availableWorkUnits)
		return nil
	}

	return fmt.Errorf("not able to add task. require %d work units, %d available", t.unitSize, w.availableWorkUnits)
}
