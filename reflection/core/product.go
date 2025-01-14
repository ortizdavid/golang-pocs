package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Product struct {
	Id       int16 
	Name     string 
	Price    float32
	Quantity int16
}

func (p *Product) ShowData(format string) {
	switch format {
	case "string":
		fmt.Println(p.ToString())
	case "json":
		fmt.Println(p.toJson())
	case "map":
		fmt.Println(p.toMap())
	case "xml":
		fmt.Println(p.toXml())
	}
}

func (p *Product) toMap() map[string]interface{} {
	return map[string]interface{}{
		"Id": p.Id,
		"Name": p.Name,
		"Price": p.Price,
		"Quantity": p.Quantity,
	}
}

func (p *Product) toJson() string {
	result, _ := json.MarshalIndent(p, " ", "  ")
	return string(result)
}

func (p *Product) ToString() string {
	return fmt.Sprintf("%d\t%s\t%f\t%d\n", p.Id, p.Name, p.Price, p.Quantity)
}

func (p *Product) toXml() string {
	result, _ := xml.MarshalIndent(p, " ", "  ")
	return string(result)
}

func (p *Product) IsNotAvalilable() bool {
	return p.Quantity == 0
}