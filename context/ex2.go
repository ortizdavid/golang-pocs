package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	parentCtx := context.Background()

	childCtx, cancel := context.WithCancel(parentCtx)

	go doWork(childCtx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}

func doWork(ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}