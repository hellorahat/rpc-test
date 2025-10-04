package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc-test/rpcmodel"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}

	args := rpcmodel.Args{A: 5, B: 7}
	reply := rpcmodel.Reply{}
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Println("Result:", reply.Result)
}
