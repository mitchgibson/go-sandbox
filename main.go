package main

import (
	"context"
	"factory/worker"
	"fmt"
	"math/rand"
	"time"
)

type Widget struct {
	qty int32
}

func main() {
	fmt.Println("Hello World")

	ctx, _ := context.WithCancel(context.Background())

	fmt.Println("Main:Creating widget worker")
	workerOpts := worker.WorkerOptions[Widget]{
		Interval:   100 * time.Millisecond,
		BufferSize: 1,
		Action: func() Widget {
			qty := rand.Int31n(10)
			fmt.Printf("Main:Creating %d widget(s)\n", qty)
			return Widget{qty}
		},
	}

	widgetCh := worker.Worker(ctx, workerOpts)

	go func() {
		for w := range widgetCh {
			time.Sleep(1 * time.Second)
			fmt.Println("Main:Goroutine:Widget created", w.qty)
		}
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("?? Cancelling context!!!")
	ctx.Done()
}
