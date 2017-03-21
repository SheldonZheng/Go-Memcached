package main

import (
	"cache"
	"fmt"
)

func main() {
	cache.Set("test","111")
	fmt.Println(cache.Get("test"))
	/*
	go grpc.ServerStart()
	grpc.ClientTest()*/
}
