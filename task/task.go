package task

import "context"

// Task main optimizer task
type Task interface {
	Run(ctx context.Context) error
}

func AllTasks() []Task {
	ts := make([]Task, 0)
	ts = append(ts, NewTaskMakeNetworkPrivate())
	ts = append(ts, NewTaskDisableEEE())

	return ts
}
