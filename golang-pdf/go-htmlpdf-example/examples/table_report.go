package examples

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-htmlpdf-example/generator"
)



func TableReportHandler(ctx *fiber.Ctx) error {
	var pdfGen generator.HtmlPdfGenenerator
	productList := getProductList()

	data := map[string]interface{}{
		"Title":       "Product List Report",
		"ProductList": productList,
		"Count":       len(productList),
	}

	pdfBytes, err := pdfGen.GeneratePDF("templates/table-report.html", data)
	if err != nil {
		return err
	}

	pdfGen.SetOutput(ctx, pdfBytes, "table-report.pdf")
	return nil
}

type Product struct {
	Id    int
	Name  string
	Price float32
}

func getProductList() []Product {
	return []Product{
		{Id: 1, Name: "Eraser", Price: 10},
		{Id: 2, Name: "Pen", Price: 100},
		{Id: 3, Name: "Pencil", Price: 50},
		{Id: 4, Name: "NoteBook", Price: 200},
		{Id: 5, Name: "Book", Price: 300},
		{Id: 6, Name: "Paperclip", Price: 30},
	}
}
