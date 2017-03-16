package main

import "grpc"

func main() {
	go grpc.ServerStart()
	grpc.ClientTest()
}
