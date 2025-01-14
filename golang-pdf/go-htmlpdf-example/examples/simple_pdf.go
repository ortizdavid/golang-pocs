package examples

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-htmlpdf-example/generator"
)

func SimplePdfHandler(ctx *fiber.Ctx) error {
	var pdfGen generator.HtmlPdfGenenerator

	data := map[string]interface{}{
		"Title": "Simple PDF",
	}

	pdfBytes, err := pdfGen.GeneratePDF("templates/simple-pdf.html", data)
	if err != nil {
		return err
	}

	pdfGen.SetOutput(ctx, pdfBytes, "simple.pdf")
	return nil
}
