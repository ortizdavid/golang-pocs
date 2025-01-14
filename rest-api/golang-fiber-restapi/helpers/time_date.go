package helpers

import "time"

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func StringToDate(strDate string) time.Time {
	date, _ := time.Parse("2006-01-02", strDate)
	return date
}

func StringToDateTime(strDateTime string) time.Time {
	dateTime, _ := time.Parse("2006-01-02 15:04:05", strDateTime)
	return dateTime
}