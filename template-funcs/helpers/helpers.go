package helpers

import (
	"fmt"
	"strings"
)

func ConcatStrings(stringArr ...string) string {
	var result string
	for _, s := range stringArr {
		result += s
	}
	return result
}

func FormatGender(gender string) string {
	var result string
	switch strings.ToUpper(gender) {
	case "M":
		result = "Male"
	case "F":
		result = "Female"
	default:
		result = "Unknown"
	}
	return result
}

func FormatMoney(value float64) string {
	formatted := fmt.Sprintf("%.2f", value) 
	parts := strings.Split(formatted, ".")
	integerPart := parts[0]
	decimalPart := parts[1]
	// Add thousand separators
	for i := len(integerPart) - 3; i > 0; i -= 3 {
		integerPart = integerPart[:i] + "," + integerPart[i:]
	}
	return integerPart + "." + decimalPart
}