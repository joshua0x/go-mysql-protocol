package protocol

import (
	"go-mysql-protocol/util"
)

type FormatDescription struct {
	BinlogVersion uint16
	MysqlVersion string
	CreateTime uint32
	EventHeaderLength byte
	EventTypeHeaderLength string
}

func DecodeFormatDescription(buf []byte) FormatDescription {
	f := new(FormatDescription)
	var c int = 0
	var t []byte
	c, f.BinlogVersion = util.ReadUB2(buf, c)
	c, t = util.ReadBytes(buf, c, 50)
	_, t = util.ReadWithNull(t, 0)
	f.MysqlVersion = string(t)
	c, f.CreateTime = util.ReadUB4(buf, c)
	c, f.EventHeaderLength = util.ReadByte(buf, c)
	c, f.EventTypeHeaderLength = util.ReadString(buf, c)
	return *f
}
