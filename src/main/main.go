package main

import (
	"net/http"
	"controller"
)

func main() {
	/*cache.Set("test","111")
	fmt.Println(cache.Get("test"))*/
	/*go grpc.ServerStart()
	grpc.ClientTest()*/
	/*messages := make(chan string,10)
	go chanTest(messages)
	for i := 0;i < 10 ;i ++ {
		fmt.Println(<-messages)
	}*/


	/*if os.Args[1] == "server" {
		tcp.StartTCPServer("9999")
	} else if os.Args[1] == "client" {
		tcp.StartClient("127.0.0.1:9999")
	}*/

	http.HandleFunc("/execute",controller.Execute)
	http.ListenAndServe(":9999",nil);
	select {}
}

/*
func chanTest(messages chan string)  {
	for i := 0;i < 10 ;i ++ {
		messages <- "str" + strconv.Itoa(i)
	}
}
*/
