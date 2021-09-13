package optimizer

import "context"

// Task main optimizer task
type Task interface {
	Run(ctx context.Context) error
}

type optimizer struct {
	tasks []Task
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

func NewOptimizer(ts []Task) Optimizer {
	return &optimizer{
		tasks: ts,
	}
}
