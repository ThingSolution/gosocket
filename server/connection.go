package server

import "net"

type Connection struct {
	net.Conn
	Id string
	RealConnection net.Conn
}

func NewConnection(connection net.Conn) *Connection {
	return &Connection{
		Id:"sdsds",
		RealConnection:connection,
	}
}
