package taskscheduler

import (
	"fmt"

	"github.com/google/uuid"
)

// Scheduler schedule units of work to Worker nodes
type Scheduler struct {
	id      string
	workers []*Worker
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		id:      uuid.New().String(),
		workers: []*Worker{},
	}
}

// AddWorker will add a new Worker to the Scheduler
//
// Trying to add a duplicate worker will result in an error
func (s *Scheduler) AddWorker(w *Worker) error {
	if containsWorker(s, w) {
		return fmt.Errorf("scheduler: %s already contains worker: %s", s.id, w.id)
	}
	s.workers = append(s.workers, w)
	return nil
}

func containsWorker(s *Scheduler, w *Worker) bool {
	for _, sw := range s.workers {
		if sw.id == w.id {
			return true
		}
	}
	return false
}

// AddTask will create a new task to be scheduled on a worker
func (s *Scheduler) AddTask(t *Task) error {
	return nil
}
