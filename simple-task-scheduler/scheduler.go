package taskscheduler

// Scheduler schedule units of work to Worker nodes
type Scheduler struct {
	workers []Worker
}
