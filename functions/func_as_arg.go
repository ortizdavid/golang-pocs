package main

import (
	"errors"
	"fmt"
)

// ProcessVariadicFn executes the given function with variadic arguments.
func ProcessVariadicFn(fn func(args ...interface{}), args ...interface{}) {
	fmt.Println("\nExecuting variadic function")
	fn(args...)
}

// ProcessSingleArgFn executes the given function with a single argument.
func ProcessSingleArgFn(fn func(arg interface{}) interface{}, arg interface{}) {
	fmt.Println("\nExecuting single argument function")
	fn(arg)
}

// ProcessSingleArgFn executes the given function with a single argument.
func ProcessSingleArgFn2[T any, R any](fn func(arg T) R, arg T) {
	fmt.Println("Executing single argument function")
	result := fn(arg)
	fmt.Println("Result:", result)
}

func main() {
	printArguments := func(args ...interface{}) {
		fmt.Println("Arguments:", args)
	}

	obj := struct {
		name string
		age int
	}{
		name: "hello",
		age: 12,
	}

	inspectObj := func(a interface{}) (any, error) {
		fmt.Println(a)

		err := errors.New("Nothing")
		return nil, err
	}

	ProcessVariadicFn(printArguments, "Hello", 42, true)
	ProcessSingleArgFn2(inspectObj, obj)
}
