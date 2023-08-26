package async

import "context"

type Future interface {
	Await() any
	AwaitWithContext(ctx context.Context) any
}

type future struct {
	ctx   context.Context
	await func(ctx context.Context) any
}

func (f future) Await() any {
	return f.await(context.Background())
}

func (f future) AwaitWithContext(ctx context.Context) any {
	return f.await(ctx)
}

func Run(f func() any) Future {
	return RunWithContext(context.Background(), f)
}

func RunWithContext(ctx context.Context, f func() any) Future {
	var result any
	c := make(chan struct{})
	go func() {
		defer close(c)
		result = f()
	}()
	return future{
		ctx: ctx,
		await: func(ctx context.Context) any {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return result
			}
		},
	}
}
