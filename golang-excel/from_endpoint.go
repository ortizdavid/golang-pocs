package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tealeg/xlsx"
)

func generateExcel() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}

	//Title
	titleRow := sheet.AddRow()
	titleCell := titleRow.AddCell()
	titleCell.Value = "Employee Information"
	//titleCell.Merge(2, 0)
	titleStyle := xlsx.NewStyle()
	titleStyle.Font = *xlsx.NewFont(16, "Arial")
	titleStyle.Font.Bold = true
	titleCell.SetStyle(titleStyle)

	// Table header
	header := sheet.AddRow()
	header.AddCell().SetString("Name")
	header.AddCell().SetString("Age")
	header.AddCell().SetString("City")

	// Set header row background color
	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "FFA500", "FFA500") 
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

	return file, nil
}

func excelHandler(c *fiber.Ctx) error {
	excelFile, err := generateExcel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating Excel file")
	}

	c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=output.xlsx")

	err = excelFile.Write(c.Response().BodyWriter())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error writing Excel file")
	}

	return nil
}

func main() {
	app := fiber.New()

	// Endpoint to generate and download Excel file
	app.Get("/", excelHandler)

	// Start the Fiber app
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting Fiber server:", err)
	}
}
