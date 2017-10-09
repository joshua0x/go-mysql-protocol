package protocol

import (
	"go-mysql-protocol/util"
	"net"
)

type HandsharkProtocol struct {
	BodySize uint32
	Sequence byte

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

func ReadHandshark(conn net.Conn) HandsharkProtocol {
	hs := new(HandsharkProtocol)
	hs.BodySize = util.ReadUB3(conn)
	hs.Sequence = util.ReadByte(conn)

	hs.ProtocolVersion = util.ReadByte(conn)
	hs.ServerVersion = string(util.ReadWithNull(conn))
	hs.ServerThreadID = util.ReadUB4(conn)
	hs.Seed = util.ReadWithNull(conn)



	hs.ServerCapabilitiesLow = util.ReadUB2(conn)
	hs.CharSet = util.ReadByte(conn)
	hs.ServerStatus = util.ReadUB2(conn)

	hs.ServerCapabilitiesHeight = util.ReadUB2(conn)


	util.ReadBytes(conn, 11)

	hs.RestOfScrambleBuff = util.ReadWithNull(conn)

	hs.Auth_plugin_name = string(util.ReadWithNull(conn))

	return *hs
}
