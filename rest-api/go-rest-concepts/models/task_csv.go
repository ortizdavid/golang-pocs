package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-rest-concepts/entities"
)

const (
	userIdIndex			= 0
	taskNameIndex		= 1
	startDateIndex		= 2
	endDateIndex		= 3
	descriptionIndex	= 4
)

func ParseTaskFromCSV(reader *csv.Reader) ([]entities.Task, error) {
	var tasks []entities.Task
	var lineNumber int

	for {
		record, err := reader.Read()
		lineNumber++
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading CSV record %v on line %d", err, lineNumber)
		}
		if len(record) != 5 {
			return nil, fmt.Errorf("invalid CSV record %v, requires: UserId,TaskName,StartDate,EndDate,Description", err)
		}
		task := entities.Task {
			TaskId:       0,
			UserId:       conversion.StringToInt(record[userIdIndex]),
			TaskName:     record[taskNameIndex],
			StartDate:    record[startDateIndex],
			EndDate:      record[endDateIndex],
			Description:  record[descriptionIndex],
			UniqueId:     encryption.GenerateUUID(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func SkipCsvHeader(reader *csv.Reader) error {
	_, err := reader.Read()
	return err
}