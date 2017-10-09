package socket

import (
	"fmt"
	"net"
	"os"
)

func GetSocket(host string, port int) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error: %s", err)
		os.Exit(1)
	}

	return conn

}

func ReadPacket(conn net.Conn) []byte {
	buf := make([]byte, 100)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("\n\n====ReadContentRaw[%d]====\n[%s]\n====EndReadContent====\n", n, string(buf))
	fmt.Printf("\n\n====ReadContent16Hex[%d]====\n[%x]\n====EndReadContent====\n", n, buf)
	return buf[0:n]
}

func WritePacket(conn net.Conn, buf []byte) {
	conn.Write(buf)
	fmt.Printf("\n\n====SendContentRaw[%d]====\n[%s]\n====EndSendContent====\n", len(buf), string(buf))
	fmt.Printf("\n\n====SendContent16Hex[%d]====\n[%x]\n====EndSendContent====\n", len(buf), buf)
}
