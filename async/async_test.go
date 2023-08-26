package async_test

import (
	"context"
	"testing"
	"time"

	"github.com/arielsrv/go-await/async"
	"github.com/stretchr/testify/assert"
)

func TestFuture_Await(t *testing.T) {
	f := func() int {
		time.Sleep(1 * time.Millisecond)
		return 1
	}

	future := async.Run(func() any {
		return f()
	})

	actual := future.Await()
	assert.Equal(t, 1, actual)
}

func TestFuture_AwaitWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	f := func() int {
		time.Sleep(1 * time.Millisecond)
		return 1
	}

	future := async.RunWithContext(ctx, func() any {
		return f()
	})

	actual := future.AwaitWithContext(ctx)
	assert.Equal(t, 1, actual)
}

func TestFuture_AwaitWithContext_Err(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	f := func() int {
		time.Sleep(10 * time.Millisecond)
		return 1
	}

	future := async.RunWithContext(ctx, func() any {
		return f()
	})

	actual := future.AwaitWithContext(ctx)
	assert.Error(t, context.DeadlineExceeded, actual)
}
