package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{-34, -3, 59, 0, 13, 1, 2, 3, 4}
	slices.Sort(s1)
	///fmt.Println(s1)
	fmt.Println(slices.Contains(s1, -34))

	s2 := []int{7, 8, 9, 12, 5, 28}
	s3 := []int{-3, 10, 5}
	s4 := slices.Concat(s1, s2, s3)
	slices.Sort(s4)
	fmt.Println(s4)
}
