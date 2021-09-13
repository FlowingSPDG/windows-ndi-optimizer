package optimizer

import (
	"context"

	"github.com/FlowingSPDG/windows-ndi-optimizer/task"
)

type optimizer struct {
	tasks []task.Task
}

func (o *optimizer) Run(ctx context.Context) error {
	for _, t := range o.tasks {
		if err := t.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}

type Optimizer interface {
	Run(ctx context.Context) error
}

func NewOptimizer(ts []task.Task) Optimizer {
	return &optimizer{
		tasks: ts,
	}
}
