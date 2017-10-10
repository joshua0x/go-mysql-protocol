package protocol

import (
	"go-mysql-protocol/util"
)

type EventFormatDescriptionBody struct {
	BinlogVersion uint16
	MysqlVersion string
	CreateTime uint32
	EventHeaderLength byte
	EventTypeHeaderLength string
}

type EventFormatDescription struct {
	Header EventHeader
	Body EventFormatDescriptionBody
}

func DecodeFormatDescription(buf []byte) EventFormatDescription {
	f := new(EventFormatDescription)
	f.Header = DecodeEventHeader(buf)

	var c int = 20
	var t []byte
	c, f.Body.BinlogVersion = util.ReadUB2(buf, c)
	c, t = util.ReadBytes(buf, c, 50)
	_, t = util.ReadWithNull(t, 0)
	f.Body.MysqlVersion = string(t)
	c, f.Body.CreateTime = util.ReadUB4(buf, c)
	c, f.Body.EventHeaderLength = util.ReadByte(buf, c)
	c, f.Body.EventTypeHeaderLength = util.ReadString(buf, c)
	return *f
}
