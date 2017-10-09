package protocol

import "go-mysql-protocol/util"

type Error struct {
	PacketType byte
	ErrorNum uint16
	ServerStatusSign byte
	ServerStatus []byte
	ServerMsg string
}

func DecodeError(buf []byte) Error {
	var cursor int
	e := new(Error)
	cursor, e.PacketType = util.ReadByte(buf, 0)
	cursor, e.ErrorNum = util.ReadUB2(buf, cursor)
	cursor, e.ServerStatusSign = util.ReadByte(buf, cursor)
	cursor, e.ServerStatus = util.ReadBytes(buf, cursor, 5)
	_, ret := util.ReadBytes(buf, cursor, len(buf) - cursor)
	e.ServerMsg = string(ret)
	return *e
}
