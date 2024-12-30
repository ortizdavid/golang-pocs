package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	
	go func() {
		time.Sleep(time.Second)
		ch <- 10
	}()
	data := <-ch
	
	fmt.Println(data)

}