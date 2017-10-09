package protocol

import (
	"go-mysql-protocol/util"
	"go-mysql-protocol/common"
)

func GetCapabilities(hs HandsharkProtocol) uint32 {
	var capabilities uint32 = 0
	capabilities |= common.CLIENT_LONG_PASSWORD
	capabilities |= common.CLIENT_FOUND_ROWS
	capabilities |= common.CLIENT_LONG_FLAG
	capabilities |= common.CLIENT_CONNECT_WITH_DB
	capabilities |= common.CLIENT_ODBC
	capabilities |= common.CLIENT_IGNORE_SPACE
	capabilities |= common.CLIENT_PROTOCOL_41
	capabilities |= common.CLIENT_INTERACTIVE
	capabilities |= common.CLIENT_IGNORE_SIGPIPE
	capabilities |= common.CLIENT_TRANSACTIONS
	capabilities |= common.CLIENT_SECURE_CONNECTION

	return capabilities
}

/**
 * 生成登录验证报文
 */
func EncodeLogin(hs HandsharkProtocol, uname string, password string, dbname string) []byte {
	buf := []byte{}

	capabilities := GetCapabilities(hs)
	capabilities |= common.CLIENT_CONNECT_WITH_DB

	buf = util.WriteUB4(buf, capabilities)
	buf = util.WriteUB4(buf, 1024 * 1024 * 16)
	buf = util.WriteByte(buf, hs.CharSet)
	for i := 0; i < 23 ; i++  {
		buf = append(buf, 0)
	}
	if len(uname) == 0 {
		buf = append(buf, 0)
	} else {
		buf = util.WriteWithNull(buf, []byte(uname))
	}

	encryPass := util.GetPassword([]byte(password), hs.Seed, hs.RestOfScrambleBuff)
	if (capabilities & common.CLIENT_SECURE_CONNECTION) > 0 {
		buf = util.WriteWithLength(buf, encryPass)
	} else {
		buf = util.WriteBytes(buf, encryPass)
		buf = util.WriteByte(buf, 0)
	}

	buf = util.WriteWithNull(buf, []byte(dbname))
	buf = util.WriteWithNull(buf, []byte(hs.Auth_plugin_name))

	return buf
}
