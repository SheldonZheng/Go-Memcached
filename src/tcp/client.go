package tcp

import (
	"net"
	"fmt"
	"os"
	"util"
	"bufio"
	"strings"
)

func StartClient(tcpaddr string){

	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	util.CheckError(err,"ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err,"DialTCP")
	//启动客户端发送线程
	go chatSend(conn)

	buf := make([]byte,1024)
	for{

		lenght, err := conn.Read(buf)
		if(util.CheckError(err,"Connection")==false){
			conn.Close()
			fmt.Println("Server is dead ...ByeBye")
			os.Exit(0)
		}
		fmt.Println(string(buf[0:lenght]))

	}
}


func chatSend(conn net.Conn){

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your command:")
	//username := conn.LocalAddr().String()
	for {

		input, err := inputReader.ReadString('\n')
		input = strings.Replace(input,"\n","",-1)
		if input == "/quit"{
			fmt.Println("ByeBye..")
			conn.Close()
			os.Exit(0);
		}


		_ ,err =conn.Write([]byte(input))
		//fmt.Println(lens)
		if(err != nil){
			fmt.Println(err.Error())
			conn.Close()
			break
		}

	}

}