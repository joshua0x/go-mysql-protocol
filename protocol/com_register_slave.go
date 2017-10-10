package protocol

import (
	"net"
	"go-mysql-protocol/util"
)

func EncodeRegisterSlave(conn net.Conn, dbIP string, dbPort int, dbName string, dbPass string, serverID uint64) []byte {
	buf := []byte{}

	buf = append(buf, 0x15)

	//从服务器ID
	buf = append(buf, byte(serverID & 0xFF))
	buf = append(buf, byte((serverID >> 8) & 0xFF))
	buf = append(buf, byte((serverID >> 16) & 0xFF))
	buf = append(buf, byte((serverID >> 24) & 0xFF))

	//主服务器IP
	buf = util.WriteLength(buf, int64(len(dbIP)))
	buf = util.WriteBytes(buf, []byte(dbIP))

	//主服务器用户名
	buf = util.WriteLength(buf, int64(len(dbName)))
	buf = util.WriteBytes(buf, []byte(dbName))

	//主服务器密码
	buf = util.WriteLength(buf, int64(len(dbPass)))
	buf = util.WriteBytes(buf, []byte(dbPass))

	//端口
	buf = util.WriteUB2(buf, uint16(dbPort))

	//安全备份级别
	buf = util.WriteByte(buf, 0x00)
	buf = util.WriteByte(buf, 0x00)
	buf = util.WriteByte(buf, 0x00)
	buf = util.WriteByte(buf, 0x00)

	//主服务ID，恒为0
	buf = util.WriteByte(buf, 0x00)
	buf = util.WriteByte(buf, 0x00)
	buf = util.WriteByte(buf, 0x00)
	buf = util.WriteByte(buf, 0x00)

	return buf
}
