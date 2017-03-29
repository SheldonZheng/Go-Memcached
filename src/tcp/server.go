package tcp

import (
	"net"
	"util"
	"fmt"
	"strings"
	"cache"
)

func StartTCPServer(port string)  {
	service := ":" + port
	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)
	util.CheckError(err,"ResolveTCPAddr")
	l,err := net.ListenTCP("tcp",tcpAddr)
	util.CheckError(err,"ListenTCP")
	conns := make(map[string]net.Conn)
	messages := make(chan string,10)
	go echoHandler(&conns,messages)

	for {
		fmt.Println("TCP Server Listening ...")
		conn,err := l.Accept()
		util.CheckError(err,"Accept")
		fmt.Println("TCP Server Accepting ...")
		conns[conn.RemoteAddr().String()]=conn
		go Handler(conn,messages)
	}
}

func Handler(conn net.Conn,messages chan string){

	fmt.Println("connection is connected from ...",conn.RemoteAddr().String())

	buf := make([]byte,1024)
	for{
		lenght, err := conn.Read(buf)
		if(util.CheckError(err,"Connection")==false){
			conn.Close()
			break
		}
		if lenght > 0{
			buf[lenght]=0
		}
		reciveStr :=string(buf[0:lenght])
		fmt.Println("Rec[",conn.RemoteAddr().String(),"] Command :" ,reciveStr)
		messages <- reciveStr
		args := strings.Split(reciveStr," ")
		if len(args) == 2 && args[0] == "get" {
			value := cache.Get(args[1])
			conn.Write([]byte(value))
		} else if len(args) == 3 && args[0] == "set" {
			cache.Set(args[1],args[2])
			conn.Write([]byte("success set key : " + args[1]))
		}

	}

}

func echoHandler(conns *map[string]net.Conn,messages chan string){


	/*for{
		msg:= <- messages
		fmt.Println(msg)
		for key,value := range *conns {
			fmt.Println("connection is connected from ...",key)
			_,err :=value.Write([]byte(msg))
			if(err != nil){
				fmt.Println(err.Error())
				delete(*conns,key)
			}

		}
	}*/

}