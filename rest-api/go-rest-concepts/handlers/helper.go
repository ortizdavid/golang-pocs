package handlers

import (
	"net/http"

	"github.com/ortizdavid/go-nopain/conversion"
)

func GetCurrentPageAndLimit(r *http.Request) (int, int) {
	currentStr := r.URL.Query().Get("current_page")
	limitStr := r.URL.Query().Get("limit")
	if currentStr == "" { currentStr = "0" }
	if limitStr == "" { limitStr = "10" }

	currentPage := conversion.StringToInt(currentStr)
	limit := conversion.StringToInt(limitStr)
	return currentPage, limit
}