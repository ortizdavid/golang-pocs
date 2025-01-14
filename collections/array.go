package main

import "fmt"

func oddNumbers(array []int) []int {
	var odds []int
	for _, elem := range array {
		if elem %2 != 0 {
			odds = append(odds, elem)
		}
	}
	return odds
}

func evenNumbers(array []int) []int{
	var evens []int
	for _, elem := range array {
		if elem %2 == 0 {
			evens = append(evens, elem)
		}
	}
	return evens
}

func printArray(array []int) {
	fmt.Print("[")
	for _, elem := range array {
		fmt.Printf("%d ", elem)
	}
	fmt.Print("]")
}

func main() {

	var array [5]int

	fmt.Println("Enter elements. ")
	for i:=0; i<len(array); i++ {
		fmt.Printf("Element V[%d]: ", i)
		fmt.Scanln(&array[i])
	}

	fmt.Print("\nOdds: ")
	printArray(oddNumbers(array[:]))

	fmt.Print("\nEvens: ")
	printArray(evenNumbers(array[:]))
}