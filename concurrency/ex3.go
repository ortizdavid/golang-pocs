package main

import (
	"fmt"
	"time"
)

func printNumber(num int)  {
	for i := 0; i < num; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	start := time.Now()

	printNumber(5)
	printNumber(3)
	printNumber(10)

	fmt.Println("Time elapsed: ", time.Since(start))

}