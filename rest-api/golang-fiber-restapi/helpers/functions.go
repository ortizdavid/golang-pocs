package helpers

import (
	"strconv"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uniqueId := uuid.New()
	return uniqueId.String()
}

func ConvertToInt(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}