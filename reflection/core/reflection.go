package core

import (
	"fmt"
	"reflect"
)

type Reflection struct {}

func (ref *Reflection) getFields(obj interface{}) []string {
	objType := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < objType.NumField(); i++ {
		fields = append(fields, objType.Field(i).Name)
	}
	return fields
}

func (ref *Reflection) getPublicMethods(obj interface{}) []string {
	var objType reflect.Type
	switch reflect.TypeOf(obj).Name() {
	case "Polygon":
		objType = reflect.TypeOf(&Polygon{})
	case "Product":
		objType = reflect.TypeOf(&Product{})
	default:
		objType = reflect.TypeOf(obj)
	}
	var methods []string
	for i := 0; i < objType.NumMethod(); i++ {
		methods = append(methods, objType.Method(i).Name)
	}
	return methods
}


func (ref *Reflection) ExistsField(obj interface{}, fieldName string) bool {
	field := reflect.ValueOf(obj).FieldByName(fieldName)
	return field.IsValid() 
}

func (ref *Reflection) ExistsMethod(obj interface{}, methodName string) bool {
	method := reflect.ValueOf(obj).MethodByName(methodName)
	return method.IsValid() 
}
	
func (ref *Reflection) CallMethod(any interface{}, methodName string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(methodName).Call(inputs)
}

func (ref *Reflection) Inspect(obj interface{})  {
	fmt.Print("\nKind: ", reflect.TypeOf(obj).Kind())
	fmt.Print("\nType: ", reflect.TypeOf(obj))
	fmt.Print("\nType Name: ", reflect.TypeOf(obj).Name())
	fmt.Print("\nAttributes: ", ref.getFields(obj))
	fmt.Print("\nPublic Methods: ", ref.getPublicMethods(obj))
	fmt.Println()
}

func (ref *Reflection) InspectFields(obj interface{})  {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	numFields := objType.NumField()
	fmt.Println("Num. of fields: ", numFields)
	fmt.Println("Field\tValue\tKind")
	for i := 0; i < numFields; i++ {
		fmt.Printf("%s\t%#v\t%v\n", objType.Field(i).Name, objVal.Field(i), objVal.Field(i).Kind())
	}
}

func (ref *Reflection) CountPrimitiveTypes(elements ...interface{}) {
	
	countFloats, countInts, countStrings := 0, 0, 0
	countBools, countMaps := 0, 0
	
	for _, elem := range elements {
		switch elem.(type) {
		case float32:	
		case float64:
			countFloats++
		case int:
			countInts++
		case string:
			countStrings++
		case bool:
			countBools++
		case map[string]interface{}:
			countMaps++
		}
	}
	fmt.Println("Integers: ", countInts)
	fmt.Println("Floats: ", countFloats)
	fmt.Println("Strings: ", countStrings)
	fmt.Println("Booleans: ", countBools)
	fmt.Println("Maps: ", countMaps)
}