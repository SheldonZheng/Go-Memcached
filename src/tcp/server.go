package tcp

import (
	"net"
	"util"
	"fmt"
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
		fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		reciveStr :=string(buf[0:lenght])
		messages <- reciveStr

	}

}

func echoHandler(conns *map[string]net.Conn,messages chan string){


	for{
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
	}

}