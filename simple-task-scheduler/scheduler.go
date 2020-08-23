package taskscheduler

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-hclog"

	"github.com/google/uuid"
)

// SchedulerOptions provides overridable options
type SchedulerOptions struct {
	ID     string
	Logger hclog.Logger
}

// Scheduler schedule units of work to Worker nodes
type Scheduler struct {
	id      string
	workers []*Worker
	logger  hclog.Logger
}

func NewScheduler(options *SchedulerOptions) *Scheduler {
	if options == nil {
		options = &SchedulerOptions{
			ID:     uuid.New().String(),
			Logger: hclog.NewNullLogger(),
		}
	}

	if options.ID == "" {
		options.ID = uuid.New().String()
	}

	return &Scheduler{
		id:      options.ID,
		workers: []*Worker{},
		logger:  options.Logger,
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
	s.logger.Info("worker added", "scheduler_id", s.id, "worker_id", w.id)
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
	for _, w := range s.workers {
		err := w.AddTask(t)

		if err != nil {
			s.logger.Info("not able to add task to worker", "worker_id", w.id, "error", err)
			continue
		}

		s.logger.Info("task added to worker", "task_id", t.id, "worker_id", w.id)
		return nil
	}

	s.logger.Error("no workers available to schedule task", "task_id", t.id)
	return errors.New("no workers available")
}
