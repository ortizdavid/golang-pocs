package main
/**
import (
	"fmt"
	"github.com/tealeg/xlsx"
)

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

// SaveToFile saves the Excel file to the specified filename.
func (eg *ExcelGenerator) SaveToFile(filename string) error {
	return eg.File.Save(filename)
}

func main() {
	// Create a new ExcelGenerator instance
	excelGenerator := NewExcelGenerator()

	// Add title, header, and data rows
	excelGenerator.AddTitle("Employee Information")
	excelGenerator.AddHeaderRow("Name", "Age", "City")
	excelGenerator.AddDataRow("Alice", 30, "New York")
	excelGenerator.AddDataRow("Bob", 35, "Los Angeles")
	excelGenerator.AddDataRow("Bob", 35, "Los Angeles")
	excelGenerator.AddDataRow("Bob", 35, "Los Angeles")
	excelGenerator.AddDataRow("Bob", 35, "Los Angeles")
	excelGenerator.AddDataRow("Bob", 35, "Los Angeles")
	excelGenerator.AddDataRow("Bob", 35, "Los Angeles")

	// Save the Excel file
	err := excelGenerator.SaveToFile("output.xlsx")
	if err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	fmt.Println("Data written to Excel successfully.")
}
*/