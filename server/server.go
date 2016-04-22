package server

import (
	"net"
	"fmt"
	"time"
)

func NewServer(address string, port string) {
	listener, err := net.Listen("tcp", address + ":" + port)

	if err != nil {
		fmt.Println("error:" + err.Error())
	}

	fmt.Println("Server started at "+address+":"+port)


	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error:" + err.Error())
		}
		// handler connection
		go handlerConnection(conn)
	}
}

func handlerConnection(connection net.Conn) {
	daytime := time.Now().String()
	connection.Write([]byte(daytime)) // don't care about return value
}

