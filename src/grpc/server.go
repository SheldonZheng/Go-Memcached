package grpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"cache"
)

type Echo int

type Args struct {
	Key   string
	Value string
}

func (t *Echo) Hi(args string, reply *string) error {
	*reply = "test:" + args
	return nil
}


func (t *Echo) Get(args Args,reply *string) error {
	*reply = cache.Get(args.Key)
	return nil
}

func (t *Echo) Set(args Args,reply *string) error {
	cache.Set(args.Key,args.Value)
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
