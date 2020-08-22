package main

import (
	"fmt"
	"log"

	taskscheduler "github.com/bcatubig/go-projects/simple-task-scheduler"
)

func main() {

	// Create a scheduler
	s1 := taskscheduler.NewScheduler()

	// Create a worker with a max capacity of 10 work units
	w1 := taskscheduler.NewWorker(10)

	// Add worker to Scheduler
	err := s1.AddWorker(w1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s1)

}
