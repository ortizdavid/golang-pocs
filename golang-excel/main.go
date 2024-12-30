package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func writeToExcel() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Table header
	header := sheet.AddRow()
	header.AddCell().SetString("Name")
	header.AddCell().SetString("Age")
	header.AddCell().SetString("City")

	// Set header row background color to brown
	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "8B4513", "8B4513") // Brown color code: 8B4513
	style.Fill = fill

	for _, cell := range header.Cells {
		cell.SetStyle(style)
	}

	// Data rows
	data := [][]string{
		{"Alice", "30", "New York"},
		{"Bob", "35", "Los Angeles"},
	}

	// Write data to the sheet
	for _, row := range data {
		excelRow := sheet.AddRow()
		for _, cell := range row {
			excelCell := excelRow.AddCell()
			excelCell.Value = cell
		}
	}

	// Save the Excel file
	err = file.Save("output.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data written to Excel successfully.")
}

func main() {
	writeToExcel()
}
