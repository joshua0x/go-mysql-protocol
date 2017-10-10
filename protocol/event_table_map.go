package protocol

import (
	"go-mysql-protocol/util"
)

type EventTableMapBody struct {
	TableMark uint32
	SchemaLength byte
	SchemaName string
	TableLength byte
	TableName string
	ColumnNum uint64
	ColumnType []byte
	ColumnMetaLength uint64
	ColumnMetaInfo []byte
}

type EventTableMap struct {
	Header EventHeader
	Body EventTableMapBody
}

func DecodeEventTableMap(buf []byte) EventTableMap {
	e := new(EventTableMap)
	e.Header = DecodeEventHeader(buf)
	var c int = 20

	//c, e.Body.TableMark = util.ReadUB4(buf, c)
	c += 8 //@TODO 查不到资料，这8个字节到底干啥

	c, e.Body.SchemaLength = util.ReadByte(buf, c)
	c, e.Body.SchemaName = util.ReadStringWithNull(buf, c)
	c, e.Body.TableLength = util.ReadByte(buf, c)
	c, e.Body.TableName = util.ReadStringWithNull(buf, c)
	c, e.Body.ColumnNum = util.ReadLength(buf, c)
	c, e.Body.ColumnType = util.ReadBytes(buf, c, int(e.Body.ColumnNum))
	c, e.Body.ColumnMetaLength = util.ReadLength(buf, c)
	c, e.Body.ColumnMetaInfo = util.ReadBytes(buf, c, int(e.Body.ColumnNum))

	return *e
}
