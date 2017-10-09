package protocol

import (
	"go-mysql-protocol/util"
	"fmt"
)

type HandsharkProtocol struct {
	ProtocolVersion byte
	ServerVersion   string
	ServerThreadID  uint32
	Seed  			[]byte
	ServerCapabilitiesLow   uint16
	CharSet         byte
	ServerStatus    uint16
	ServerCapabilitiesHeight   uint16
	RestOfScrambleBuff []byte
	Auth_plugin_name string
}

func DecodeHandshark(buff []byte) HandsharkProtocol {
	var cursor int
	var tmp []byte
	hs := new(HandsharkProtocol)

	cursor, hs.ProtocolVersion = util.ReadByte(buff, cursor)
	cursor, tmp = util.ReadWithNull(buff, cursor)
	hs.ServerVersion = string(tmp)
	cursor, hs.ServerThreadID = util.ReadUB4(buff, cursor)
	cursor, hs.Seed = util.ReadWithNull(buff, cursor)
	cursor, hs.ServerCapabilitiesLow = util.ReadUB2(buff, cursor)
	cursor, hs.CharSet = util.ReadByte(buff, cursor)
	cursor, hs.ServerStatus = util.ReadUB2(buff, cursor)
	cursor, hs.ServerCapabilitiesHeight = util.ReadUB2(buff, cursor)
	cursor, _ = util.ReadBytes(buff, cursor, 11)
	cursor, hs.RestOfScrambleBuff = util.ReadWithNull(buff, cursor)
	cursor, tmp = util.ReadWithNull(buff, cursor)
	hs.Auth_plugin_name = string(tmp)

	fmt.Printf("DecodeHanshark: %+v\n", hs)

	return *hs
}
