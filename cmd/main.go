package main

import (
	"context"

	"github.com/FlowingSPDG/windows-ndi-optimizer/optimizer"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	tasks := make([]optimizer.Task, 0)
	opt := optimizer.NewOptimizer(tasks)
	if err := opt.Run(ctx); err != nil {
		return nil
	}
	return nil
}
