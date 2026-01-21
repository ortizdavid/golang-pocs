package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Calculator defines methods for RPC server
type Calculator struct {}

type Args struct {
	A, B int
}

// All the exposed methods -----------------------
func (t *Calculator) Sum(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (t *Calculator) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}


func main() {
	// Register services
	rpc.Register(new(Calculator))

	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":1234")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	
	log.Println("Server is running.. Port", ":1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
