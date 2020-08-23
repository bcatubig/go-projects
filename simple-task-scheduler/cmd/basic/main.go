package main

import (
	"os"

	"github.com/hashicorp/go-hclog"

	taskscheduler "github.com/bcatubig/go-projects/simple-task-scheduler"
)

func main() {
	// Create a logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "simple-task-scheduler",
		Output: os.Stdout,
	})

	// Create a scheduler
	s1 := taskscheduler.NewScheduler(&taskscheduler.SchedulerOptions{
		Logger: logger.Named("scheduler-1"),
	})

	// Create a worker with a max capacity of 3 work units
	w1 := taskscheduler.NewWorker(3, &taskscheduler.WorkerOptions{
		Logger: logger.Named("worker-1"),
	})

	w2 := taskscheduler.NewWorker(3, &taskscheduler.WorkerOptions{
		Logger: logger.Named("worker-2"),
	})

	// Add workers to Scheduler
	err := s1.AddWorker(w1)

	if err != nil {
		logger.Error("failed to add worker", "error", err)
		os.Exit(1)
	}

	err = s1.AddWorker(w2)

	if err != nil {
		logger.Error("failed to add worker", "error", err)
		os.Exit(1)
	}

	// Create some tasks
	t1 := taskscheduler.NewTask("task 1", 1)
	t2 := taskscheduler.NewTask("task 2", 3)
	t3 := taskscheduler.NewTask("task 3", 2)
	t4 := taskscheduler.NewTask("task 4", 5)

	err = s1.AddTask(t1)

	if err != nil {
		logger.Error("failed to add task", "error", err)
		os.Exit(1)
	}

	err = s1.AddTask(t2)

	if err != nil {
		logger.Error("failed to add task", "error", err)
		os.Exit(1)
	}

	err = s1.AddTask(t3)

	if err != nil {
		logger.Error("failed to add task", "error", err)
		os.Exit(1)
	}

	err = s1.AddTask(t4)

	if err != nil {
		logger.Error("failed to add task", "error", err)
		os.Exit(1)
	}
}
