package worker

import (
	"context"
	"fmt"
	"time"
)

type WorkerOptions[T any] struct {
	Interval   time.Duration
	BufferSize int
	Action     func() T
}

func Worker[T any](ctx context.Context, opts WorkerOptions[T]) <-chan T {
	ch := make(chan T, opts.BufferSize)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker:Closing channel")
				close(ch)
				return
			default:
				fmt.Println("Worker:Running action")
				ch <- opts.Action()
				fmt.Println("Worker:Buffer size:", len(ch))
				time.Sleep(opts.Interval)
			}
		}
	}()

	return ch
}
