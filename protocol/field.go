package protocol

import "go-mysql-protocol/util"

type Field struct {
	DirName string
	DatabaseName string
	TableName string
	TablePreName string
	RowName string
	RowPreName string
	CharSet uint16
	RowLength uint32
	RowType byte
	RowSign uint16
	IntDegree byte
	DefaultValue string
}

func DecodeField(buf []byte) Field {
	f := new(Field)
	cursor, tmp := util.ReadLengthString(buf, 0)
	f.DirName = string(tmp)
	cursor, tmp = util.ReadLengthString(buf, cursor)
	f.DatabaseName = string(tmp)
	cursor, tmp = util.ReadLengthString(buf, cursor)
	f.TableName = string(tmp)
	cursor, tmp = util.ReadLengthString(buf, cursor)
	f.TablePreName = string(tmp)
	cursor, tmp = util.ReadLengthString(buf, cursor)
	f.RowName = string(tmp)
	cursor, tmp = util.ReadLengthString(buf, cursor)
	f.RowPreName = string(tmp)
	cursor, _ = util.ReadByte(buf, cursor)
	cursor, f.CharSet = util.ReadUB2(buf, cursor)
	cursor, f.RowLength = util.ReadUB4(buf, cursor)
	cursor, f.RowType = util.ReadByte(buf, cursor)
	cursor, f.RowSign = util.ReadUB2(buf, cursor)
	cursor, f.IntDegree = util.ReadByte(buf, cursor)
	cursor, _ = util.ReadUB2(buf, cursor)
	if cursor < len(buf) {
		cursor, f.DefaultValue = util.ReadLengthString(buf, cursor)
	}
	return *f
}