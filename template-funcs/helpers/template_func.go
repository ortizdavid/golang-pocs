package helpers

import "github.com/gofiber/template/html/v2"


// pass functions to al templates
func AddTemplateFunc(engine *html.Engine) {
	engine.AddFunc("ConcatStrings", ConcatStrings)
	engine.AddFunc("FormatMoney", FormatMoney)
	engine.AddFunc("FormatGender", FormatGender)
}