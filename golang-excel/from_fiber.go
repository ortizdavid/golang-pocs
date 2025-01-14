package main

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tealeg/xlsx"
)

// Employee struct representing data from the database
type Employee struct {
	ID   int
	Name string
	Age  int
	City string
}

// ExcelGenerator is a struct that generates Excel files.
type ExcelGenerator struct {
	File  *xlsx.File
	Sheet *xlsx.Sheet
}

// NewExcelGenerator creates a new ExcelGenerator instance with a new sheet.
func NewExcelGenerator() *ExcelGenerator {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	return &ExcelGenerator{
		File:  file,
		Sheet: sheet,
	}
}

// AddTitle adds a title row to the Excel sheet.
func (eg *ExcelGenerator) AddTitle(title string) {
	row := eg.Sheet.AddRow()
	cell := row.AddCell()
	cell.Value = title
	titleStyle := xlsx.NewStyle()
	titleStyle.Font = *xlsx.NewFont(12, "Arial")
	titleStyle.Font.Bold = true
	cell.SetStyle(titleStyle)
}

// AddHeaderRow adds a header row with specified columns to the Excel sheet.
func (eg *ExcelGenerator) AddHeaderRow(columns ...string) {
	row := eg.Sheet.AddRow()
	for _, col := range columns {
		cell := row.AddCell()
		cell.Value = col
	}
}

// AddDataRow adds a data row to the Excel sheet.
func (eg *ExcelGenerator) AddDataRow(data ...interface{}) {
	row := eg.Sheet.AddRow()
	for _, d := range data {
		cell := row.AddCell()
		cell.Value = fmt.Sprintf("%v", d)
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// Fetch data from the database (simulated in this example)
		employees, err := fetchDataFromDatabase()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching data from the database")
		}

		// Create a new ExcelGenerator instance
		excelGenerator := NewExcelGenerator()

		// Add title, header, and data rows
		excelGenerator.AddTitle("Employee Information")
		excelGenerator.AddHeaderRow("ID", "Name", "Age", "City")

		// Populate data from the database
		for _, employee := range employees {
			excelGenerator.AddDataRow(employee.ID, employee.Name, employee.Age, employee.City)
		}

		// Create a buffer to store the Excel data
		var buf bytes.Buffer

		// Write the Excel data to the buffer
		err = excelGenerator.File.Write(&buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error writing Excel data")
		}

		// Set response headers for Excel download
		c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Response().Header.Set("Content-Disposition", "attachment; filename=new_excel.xlsx")

		// Send the Excel data to the client
		_, err = c.Write(buf.Bytes())
		if err != nil {
			return err
		}

		return nil
	})

	// Start the Fiber app
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting Fiber server:", err)
	}
}

// fetchDataFromDatabase simulates fetching data from a database.
func fetchDataFromDatabase() ([]Employee, error) {
	// Simulate fetching data from the database.
	// In a real application, you would fetch data from your actual database.
	employees := []Employee{
		{1, "Alice", 30, "New York"},
		{2, "Bob", 35, "Los Angeles"},
		{2, "Bob", 35, "Los Angeles"},
		{2, "Bob", 35, "Los Angeles"},
		{2, "Bob", 35, "Los Angeles"},
		{2, "Bob", 35, "Los Angeles"},
		{2, "Bob", 35, "Los Angeles"},
	}

	return employees, nil
}
