package main

import (
	"context"

	"github.com/FlowingSPDG/windows-ndi-optimizer/optimizer"
	"github.com/FlowingSPDG/windows-ndi-optimizer/task"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	tasks := task.AllTasks()
	opt := optimizer.NewOptimizer(tasks)
	if err := opt.Run(ctx); err != nil {
		return nil
	}
	return nil
}
