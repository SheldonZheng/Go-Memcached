package grpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Echo int

func (t *Echo) Hi(args string, reply *string) error {
	*reply = "test:" + args
	return nil
}

func ServerStart() {
	rpc.Register(new(Echo))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("rpc listener started.")
	for {
		http.Serve(l, nil)
	}
}
