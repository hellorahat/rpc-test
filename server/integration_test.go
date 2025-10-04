package main

import (
	"net"
	"net/rpc"
	"rpc-test/rpcmodel"
	"testing"
	"time"
)

func TestRPCServerClient(t *testing.T) {
	// start listener
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal("Error starting listener:", err)
	}
	defer ln.Close()

	arith := new(Arith)
	rpc.Register(arith)
	go rpc.Accept(ln)
	time.Sleep(100 * time.Millisecond)

	// connect client
	client, err := rpc.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal("Error Dialing:", err)
	}
	defer client.Close()

	// call remote method
	args := rpcmodel.Args{A: 7, B: 5}
	reply := rpcmodel.Reply{}
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		t.Fatal("RPC error:", err)
	}

	expected := 35
	if reply.Result != expected {
		t.Errorf("Expected %d, got %d", expected, reply.Result)
	}
}
