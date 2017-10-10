package helper

import (
	"net"
	"go-mysql-protocol/protocol"
	"go-mysql-protocol/socket"
	"fmt"
)

func ExecSql(conn net.Conn, sql string) protocol.ResultSet {
	//发送命令，查看binlog信息
	tmp := protocol.EncodeQuery(sql)
	socket.WritePacket(conn, 0, tmp)

	return protocol.DecodeResultSet(conn)
}

func GetTableInfo(conn net.Conn) map[string][]string {
	var tbRS, clnRS protocol.ResultSet
	var sql string

	//存放表名称和列字段的映射
	var mapTableColumns map[string][]string
	mapTableColumns = make(map[string][]string)


	sql = "show tables"
	tbRS = ExecSql(conn, sql)
	for _, tbRowData := range tbRS.Data {
		tbName := tbRowData.Data[0]
		sql = fmt.Sprintf("DESC %s", tbName)
		clnRS = ExecSql(conn, sql)
		for _, clnRowData := range clnRS.Data  {
			clnName := clnRowData.Data[0]
			mapTableColumns[tbName] = append(mapTableColumns[tbName], clnName)
		}
	}

	return mapTableColumns
}
