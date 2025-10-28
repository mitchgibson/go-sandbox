package main

import (
	"context"
	"factory/worker"
	"fmt"
	"time"
)

type Widget struct {
	qty int32
}

func main() {
	fmt.Println("Hello World")

	ctx, _ := context.WithCancel(context.Background())

	fmt.Println("?? Creating widget worker!!!")
	workerOpts := worker.WorkerOptions[Widget]{
		Interval:   100 * time.Millisecond,
		BufferSize: 10,
		Action: func() Widget {
			fmt.Println("?? Producing widget!!!")
			return Widget{qty: 1}
		},
	}

	widgetCh := worker.Worker(ctx, workerOpts)

	go func() {
		for w := range widgetCh {
			time.Sleep(1 * time.Second)
			fmt.Println("?? Consuming widget!!!", w.qty)
		}
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("?? Cancelling context!!!")
	ctx.Done()
}
