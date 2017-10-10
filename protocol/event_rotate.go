package protocol

import "go-mysql-protocol/util"

type Rotate struct {
	Offset uint64
	BinlogFileName string
}

func DecodeRotate(buf []byte) Rotate {
	var c int = 0
	r := new(Rotate)
	c, r.Offset = util.ReadUB8(buf, c)
	c, r.BinlogFileName = util.ReadString(buf, c)
	return *r
}
