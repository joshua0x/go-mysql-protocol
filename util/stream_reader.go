package util

import (
	"net"
)

func ReadBytes(conn net.Conn, offset int) []byte {
	tmp := make([]byte, offset)
	conn.Read(tmp)
	return tmp
}

func ReadByte(conn net.Conn) byte {
	tmp := make([]byte, 1)
	conn.Read(tmp)
	return tmp[0]
}

func ReadUB2(conn net.Conn) uint16 {
	tmp := make([]byte, 2)
	conn.Read(tmp)
	i := uint16(tmp[0] & 0xFF)
	i |= uint16(tmp[1] & 0xFF) << 8
	return uint16(i)
}

func ReadUB3(conn net.Conn) uint32 {
	tmp := make([]byte, 3)
	conn.Read(tmp)
	i := uint32(tmp[0] & 0xFF)
	i |= uint32((tmp[1] & 0xFF) << 8)
	i |= uint32((tmp[2] & 0xFF) << 16)
	return i
}

func ReadUB4(conn net.Conn) uint32 {
	tmp := make([]byte, 4)
	conn.Read(tmp)
	i := uint32(tmp[0] & 0xFF)
	i |= uint32((tmp[1] & 0xFF) << 8)
	i |= uint32((tmp[2] & 0xFF) << 16)
	i |= uint32((tmp[3] & 0xFF) << 24)
	return i
}

func ReadUB8(conn net.Conn) uint64 {
	tmp := make([]byte, 4)
	conn.Read(tmp)
	i := uint64(tmp[0] & 0xFF)
	i |= uint64((tmp[1] & 0xFF) << 8)
	i |= uint64((tmp[2] & 0xFF) << 16)
	i |= uint64((tmp[3] & 0xFF) << 24)
	i |= uint64((tmp[4] & 0xFF) << 32)
	i |= uint64((tmp[5] & 0xFF) << 40)
	i |= uint64((tmp[6] & 0xFF) << 48)
	i |= uint64((tmp[7] & 0xFF) << 56)
	return i
}

func ReadLength(conn net.Conn) uint64 {
	length := ReadByte(conn)
	switch length {
	case 251:
		return 0
	case 252:
		return uint64(ReadUB2(conn))
	case 253:
		return uint64(ReadUB3(conn))
	case 254:
		return ReadUB8(conn)
	default:
		return uint64(length)

	}
}

func ReadWithNull(conn net.Conn) []byte {
	ret := []byte{}
	for ; ;  {
		tmp := make([]byte, 1)
		conn.Read(tmp)
		if tmp[0] != 0 {
			ret = append(ret, tmp[0])
		} else {
			break
		}
	}
	return ret
}