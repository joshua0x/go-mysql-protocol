package socket

import (
	"fmt"
	"net"
	"os"
	"go-mysql-protocol/util"
)

func GetSocket(host string, port int) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		util.WriteErrorLog("GetSocket fail.")
		os.Exit(1)
	}

	return conn

}

// 返回值顺序为BodySize, Sequence, Body
func ReadPacket(conn net.Conn) (uint32, byte, []byte) {
	headerBuf := make([]byte, 4)
	_, err := conn.Read(headerBuf)
	if err != nil {
		return 0, 0, nil
	}
	bodySize := (uint32)(headerBuf[2] << 16) + uint32(headerBuf[1] << 8) + uint32(headerBuf[0])
	sequence := headerBuf[3]

	bodyBuf := make([]byte, bodySize)
	_, err = conn.Read(bodyBuf)
	if err != nil {
		return 0, 0, nil
	}

	fmt.Printf("socket read: %d %d\n", headerBuf, bodyBuf)
	return bodySize, sequence, bodyBuf
}

func WritePacket(conn net.Conn, sequence byte, buf []byte) {
	l := len(buf)
	s := []byte{}
	s = append(s, byte(l & 0xFF))
	s = append(s, byte((l >> 8) & 0xFF))
	s = append(s, byte((l >> 16) & 0xFF))
	s = append(s, sequence)
	for _,v := range buf  {
		s = append(s, v)
	}
	fmt.Printf("socket send: %d\n", s)
	conn.Write(s)
}
