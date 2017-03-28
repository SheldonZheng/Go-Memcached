package main

import "tcp"

func main() {
	/*cache.Set("test","111")
	fmt.Println(cache.Get("test"))*/
	/*go grpc.ServerStart()
	grpc.ClientTest()*/
	tcp.StartTCPServer("9999")

}
