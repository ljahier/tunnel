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
	for {
		// listen for an incoming connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		err = conn.SetKeepAlive(true)
		if err != nil {
			fmt.Println("Cannot keep alive connection: ", err)
		}

		// must send conn to reverse proxy which determine to what tcp keep alived conn we need to send data

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	fmt.Println(string(buf))
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
