package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
	defer cancel()

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled: ", ctx.Err())
	}
}