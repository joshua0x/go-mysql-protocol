package protocol

import (
	"net"
	"go-mysql-protocol/socket"
	"go-mysql-protocol/util"
)

type ResultSet struct {
	RowNum uint64
	Fields []Field
	Data []RowData
}

func DecodeResultSet(conn net.Conn) ResultSet {
	var body []byte
	var resultSet ResultSet
	resultSet.Fields = make([]Field, 0)
	resultSet.Data = make([]RowData, 0)

	//reset header
	_, _, body = socket.ReadPacket(conn)
	_, resultSet.RowNum = util.ReadLength(body, 0)

	//fields
	_, _, body = socket.ReadPacket(conn)
	for ; ; {
		if body[0] == 0xFE {
			break
		}
		resultSet.Fields = append(resultSet.Fields, DecodeField(body))
		_, _, body = socket.ReadPacket(conn)
	}

	//rowdata
	_, _, body = socket.ReadPacket(conn)
	for ; ;  {
		if body[0] == 0xFE {
			break
		}
		resultSet.Data = append(resultSet.Data, DecodeRowData(body))
		_, _, body = socket.ReadPacket(conn)
	}
	return resultSet
}
