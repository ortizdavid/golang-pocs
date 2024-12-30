package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numberStr := "2.9765442222222220"
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(formatFloat(number))
}

func formatFloat(number float64) string {
	// Convert the float64 to a string without any formatting
	str := strconv.FormatFloat(number, 'f', -1, 64)

	// Trim trailing zeros and dot
	str = strings.TrimRight(str, "0")
	str = strings.TrimRight(str, ".")

	return str
}