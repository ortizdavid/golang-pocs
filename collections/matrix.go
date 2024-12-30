package main

import (
	"fmt"
	"math/rand"
	"time"
)


func readMatrix(numLines, numCols int)  {
	var matrix [][]int
	matrix = make([][]int, numLines)
    for i := range matrix {
        matrix[i] = make([]int, numCols)
    }
	
	fmt.Println("Enter matrix elements.")
	for i := 0; i < numLines; i++ {
		for j := 0; j < numCols; j++ {
			fmt.Printf("Element M[%d, %d]: ", i, j)
			fmt.Scanln(&matrix[i][j])
		}
	}
}


func printMatrixRange(matrix [][]int) {
	fmt.Println("\nPrinting matrix - Range.")
	for _, row := range matrix {
		for _, element := range row {
			fmt.Printf("%d\t", element)
		}
		fmt.Println()
	}
}

func printMatrixIndex(matrix [][]int) {
	fmt.Println("\nPrinting matrix - Index.")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}


func generateRandomMatrix(numLines int, numCols int) [][]int {
	matrix := make([][]int, numLines)
	for i := range matrix {
		matrix[i] = make([]int, numCols)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = rand.Intn(99) + i
		}
	}
	return matrix
}

func minMaxElem(matrix [][]int) (int, int) {
	var max, min int
	min, max = matrix[0][0], matrix[0][0]
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
	}
	return min, max
}

func sumAverage(matrix [][]int) (int, float32) {
	var sum int
	var average float32

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum = sum + matrix[i][j]
		}
	}
	average = float32(sum) / float32(len(matrix)  *len(matrix[0]))
	return sum, average
}


func matrixInfo(matrix [][]int) {
	min, max := minMaxElem(matrix)
	sum, average := sumAverage(matrix)
	fmt.Println("\nMin: ", min)
	fmt.Println("Max: ", max)
	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", average)
}


func main() {
	
	//readMatrix(2, 2)
	//printMatrixIndex(mat1)

	/*mat2 := [][]int{
		{1, 2, 3},
		{4, 5, -6},
	}
	printMatrixIndex(mat2)*/

	mat3 := generateRandomMatrix(5, 5)
	printMatrixIndex(mat3)
	matrixInfo(mat3)
}