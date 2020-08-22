package taskscheduler

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	UNSCHEDULED = "unscheduled"
	SCHEDULED   = "scheduled"
	RUNNING     = "running"
	STOPPED     = "stopped"
	FAILED      = "failed"
)

// Task is an individual unit of work
type Task struct {
	id       string
	name     string
	status   string
	unitSize int
}

// NewTask returns a new, unscheduled task
func NewTask(name string, unitSize int) *Task {
	return &Task{
		id:       uuid.New().String(),
		name:     name,
		status:   UNSCHEDULED,
		unitSize: unitSize,
	}
}

// Start starts a task
func (t *Task) Start() error {
	fmt.Printf("Starting task: %s <%s>", t.name, t.id)
	t.status = RUNNING
	return nil
}

// Stop stops a task
func (t *Task) Stop() error {
	fmt.Printf("Stopping task: %s <%s>", t.name, t.id)
	t.status = STOPPED
	return nil
}
