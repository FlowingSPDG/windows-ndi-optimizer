package task

import "context"

// Task main optimizer task
type Task interface {
	Run(ctx context.Context) error
}

// AllTasks Get All Optimize tasks
func AllTasks() []Task {
	ts := []Task{
		NewTaskMakeNetworkPrivate(),
		NewTaskDisableEEE(),
		NewTaskMakeNetworkPrivate(),
	}

	return ts
}
