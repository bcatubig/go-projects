package taskscheduler

// Worker executes tasks
type Worker struct {
	maxTaskCapacity int
	tasks           []Task
}
