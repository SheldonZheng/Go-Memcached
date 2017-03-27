package grpc

import (
	"fmt"
	"log"
	"net/rpc"
)

func ClientTest() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var args = "hello rpc"
	var reply string

	err = client.Call("Echo.Hi", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("Arith: %d*%d=%d\n", args, reply)

	var argsa *Args = &Args{"test1", "aaaa"}

	err = client.Call("Echo.Set",&argsa,&reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("Set:%d + %d", argsa, reply)

	err = client.Call("Echo.Get",&argsa,&reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("Get:%d + %d", argsa, reply)

}
