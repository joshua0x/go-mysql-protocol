package main

import (
	"go-mysql-protocol/socket"
	"go-mysql-protocol/protocol"
	"fmt"
)

func main() {
	conn := socket.GetSocket("127.0.0.1", 3306)
	hs := protocol.ReadHandshark(conn)
	fmt.Printf("%+v\n", hs)

	protocol.WriteLogin(conn, hs, "root", "MhxzKhl", "test")
	socket.ReadPacket(conn)
}
