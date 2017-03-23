package main

import (
	"grpc"
)

func main() {
	/*cache.Set("test","111")
	fmt.Println(cache.Get("test"))*/
	go grpc.ServerStart()
	grpc.ClientTest()
}
