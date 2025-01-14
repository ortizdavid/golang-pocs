package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int)  {
	for x := range data {
		fmt.Printf("worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	start := time.Now()

	data := make(chan int)
	numWorkers := 10_000

	for i := 0; i < numWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1_000_000_000; i++ {
		data <- i
	}

	fmt.Println("Time ellapsed:", time.Since(start))
}

