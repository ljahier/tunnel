package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const HOST = "127.0.0.1"

func Init(port int) {
	fmt.Println("Hello from server")
	ch := make(chan string)
	// get somes info from cli

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", HOST, port))
	if err != nil {
		panic("The server address is not valid")
	}

	// run tcp server
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + HOST + ":" + strconv.Itoa(port))
	go reverseProxy(ch)
	for {
		// listen for an incoming connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// must send conn to reverse proxy which must determine to which tcp conn keep alived conn we need to send data

		go handleRequest(conn, ch)
	}
}

func handleRequest(conn *net.TCPConn, ch chan string) {
	err := conn.SetKeepAlive(true)
	if err != nil {
		fmt.Println("Cannot keep alive connection: ", err)
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println(string(buf))
	ch <- string(buf)
	conn.Write([]byte("Message received"))
}
