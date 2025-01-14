package main

import (
	"fmt"
	"github.com/ortizdavid/reflection-golang/core"
)

func main() {
	//----Objects--------------
	product := core.Product{
		Id: 1,
		Name: "Computer",
		Price: 1000,
		Quantity: 2,
	}
	polygon := core.Polygon{
		Height: 100.1,
		Width: 98.3,
	}
	var reflection core.Reflection

	fmt.Println("Inspecting 1st object")
	reflection.Inspect(product)
	fmt.Println("\nInspecting Fields")
	reflection.InspectFields(product)
	fmt.Println("\nCalling Method 'ShowData' with 'Json':")
	reflection.CallMethod(&product, "ShowData", "json")

	fmt.Println("\n==============================================================================")
	fmt.Println("Inspecting 2nd object")
	reflection.Inspect(polygon)
	field := "Width"
	fmt.Printf("\nField '%s' Exists? %t\n", field, reflection.ExistsField(polygon, field))
	method := "GetName"
	fmt.Printf("Method '%s' Exists? %t\n", method, reflection.ExistsMethod(polygon, method))
	fmt.Println("\nCalling Method 'Verify':")
	reflection.CallMethod(&polygon, "Verify", 100.1, 98.3)

	fmt.Println("\n==============================================================================")
	fmt.Println("Counting Primitive Types:")
	reflection.CountPrimitiveTypes(1, 11, 34, 0.9, 1.5, "Hello", "Go", true, false, map[string]interface{}{"x": 12})
}