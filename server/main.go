package main

import (
	"fmt"
	"net"
	"net/rpc"
	"rpc-test/rpcmodel"
)

type Arith struct{}

func (t *Arith) Multiply(args rpcmodel.Args, reply *rpcmodel.Reply) error {
	reply.Result = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on port 1234")
	rpc.Accept(ln)
}
