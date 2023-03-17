package client

import (
	"fmt"
	"net"
)

func Init(server string, port int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		panic("The server address is not valid")
	}
	fmt.Println("hello")
	// get an subdomain
	// show subdomain
	// each incom from this connection must be redirect to local http app which are defined at the client run

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err.Error())
	}

	err = conn.SetKeepAlive(true)
	if err != nil {
		fmt.Println("Cannot keep alive connection: ", err)
	}

	conn.Write([]byte("Hello")) // we just need to open tcp conn
}
