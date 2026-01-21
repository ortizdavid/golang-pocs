package main

import (
    "fmt"
    "log"
    "net/rpc/jsonrpc"
)

type Args struct {
    A, B int
}

func main() {
    client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("Erro ao conectar:", err)
    }
    defer client.Close() 

    args := &Args{A: 9, B: 6}
    var reply int

    // 1. Calls Sum
    err = client.Call("Calculator.Sum", args, &reply)
    if err != nil {
        log.Fatal("Error calling sum service:", err)
    }
    fmt.Printf("Sum: %d + %d = %d\n", args.A, args.B, reply)

    // 2. Calls Multiply
    err = client.Call("Calculator.Multiply", args, &reply)
    if err != nil {
        log.Fatal("Error calling multiply service:", err)
    }
    fmt.Printf("Multiply: %d * %d = %d\n", args.A, args.B, reply)
}