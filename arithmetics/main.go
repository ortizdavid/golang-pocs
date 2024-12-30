package main

import "fmt"

func main() {

	x := 0.1 + (0.2 + 0.3)
	y := (0.1 + 0.2) + 0.3
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("x == y --> ", x==y)
}
