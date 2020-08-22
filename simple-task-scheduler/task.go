package taskscheduler

const (
	UNSCHEDULED = "unscheduled"
	SCHEDULED   = "scheduled"
	RUNNING     = "running"
	FAILED      = "failed"
)

// Task is an individual unit of work
type Task struct {
	name     string
	status   string
	unitSize int
}

// NewTask returns a new, unscheduled task
func NewTask(name string, unitSize int) *Task {
	return &Task{
		name:     name,
		status:   UNSCHEDULED,
		unitSize: unitSize,
	}
}
