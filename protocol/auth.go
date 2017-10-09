package protocol

import (
	"go-mysql-protocol/util"
	"net"
	"fmt"
	"go-mysql-protocol/common"
)

func GetCapabilities(hs HandsharkProtocol) uint32 {
	/**/
	noSchema := 1 << 4
	longPassword := 1
	longFlag := 1 << 2
	connectWithDb := 1 << 3
	transactions := 1 << 13
	secureConnection := 1 << 15
	protocol41 := 1 << 9

	return uint32(longPassword | longFlag | connectWithDb | transactions | protocol41 | secureConnection | noSchema)
	/**/
	/**
	foundRows := 1 << 1
	connectWithDb := 1 << 3
	compress := 1 << 5
	odbc := 1 << 6
	localFiles := 1 << 7
	ignoreSpace := 1 << 8
	multiStatements := 1 << 16
	multiResults := 1 << 17
	interactive := 1 << 10
	ssl := 1 << 11
	ignoreSigPipe := 1 << 12

	return int64(foundRows | connectWithDb | compress | odbc | localFiles | ignoreSpace | multiStatements | multiResults | interactive | ssl | ignoreSigPipe)
	**/
}

/**
 * 生成登录验证报文
 */
func WriteLogin(conn net.Conn, hs HandsharkProtocol, uname string, password string, dbname string) {
	buf := []byte{}
	buf = append(buf, 0)
	buf = append(buf, 0)
	buf = append(buf, 0)
	buf = append(buf, 0)

	//buf = util.WriteUB4(buf, int32((hs.ServerCapabilitiesHeight << 16) + hs.ServerCapabilitiesLow))
	//capabilities := (uint32(hs.ServerCapabilitiesHeight) << 16) | uint32(hs.ServerCapabilitiesLow)
	capabilities := GetCapabilities(hs)
	capabilities |= common.CLIENT_CONNECT_WITH_DB

	buf = util.WriteUB4(buf, capabilities)
	buf = util.WriteUB4(buf, 0)
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

	length := len(buf) - 4
	buf[0] = byte(length & 0xFF)
	buf[1] = byte((length >> 8) & 0xFF)
	buf[2] = byte((length >> 16) & 0xFF)
	buf[3] = 1

	fmt.Println("=====:", string(buf))

	fmt.Printf("%x\n", buf)

	conn.Write(buf)
}
