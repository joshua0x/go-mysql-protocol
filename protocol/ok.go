package protocol

import (
	"go-mysql-protocol/util"
)

type OK struct {
	PacketType 		byte
	AffectedRows 	uint64
	InsertID 		uint64
	ServerStatus 	uint16
	WarningNum 		uint16
	ServerMsg 		[]byte
}

func DecodeOk(buff []byte) OK {
	var cursor int
	ok := new(OK)
	cursor, ok.PacketType 	= util.ReadByte(buff, 0)
	cursor, ok.AffectedRows = util.ReadLength(buff, cursor)
	cursor, ok.InsertID 	= util.ReadLength(buff, cursor)
	cursor, ok.ServerStatus = util.ReadUB2(buff, cursor)
	cursor, ok.WarningNum 	= util.ReadUB2(buff, cursor)
	return *ok
}
