package util

import (
	"github.com/willf/bitset"
)

func ReadBytes(buff []byte, cursor int, offset int) (int, []byte) {
	if offset <= 0 {
		return cursor, nil
	}
	return cursor + offset, buff[cursor:cursor + offset]
}

func ReadByte(buff []byte, cursor int) (int, byte) {
	return cursor + 1, buff[cursor]
}

func ReadUB2(buff []byte, cursor int) (int, uint16) {
	i := uint16(buff[cursor])
	i |= uint16(buff[cursor + 1]) << 8
	return cursor + 2, i
}

func ReadUB3(buff []byte, cursor int) (int, uint32) {
	i := uint32(buff[cursor])
	i |= uint32(buff[cursor + 1]) << 8
	i |= uint32(buff[cursor + 2]) << 16
	return cursor + 3, i
}

func ReadUB4(buff []byte, cursor int) (int, uint32) {
	i := uint32(buff[cursor])
	i |= uint32(buff[cursor + 1]) << 8
	i |= uint32(buff[cursor + 2]) << 16
	i |= uint32(buff[cursor + 3]) << 24
	return cursor + 4, i
}

func ReadUB6(buff []byte, cursor int) (int, uint64) {
	i := uint64(buff[cursor])
	i |= uint64(buff[cursor + 1]) << 8
	i |= uint64(buff[cursor + 2]) << 16
	i |= uint64(buff[cursor + 3]) << 24
	i |= uint64(buff[cursor + 4]) << 32
	i |= uint64(buff[cursor + 5]) << 40
	return cursor + 6, i
}

func ReadUB8(buff []byte, cursor int) (int, uint64) {
	i := uint64(buff[cursor])
	i |= uint64(buff[cursor + 1]) << 8
	i |= uint64(buff[cursor + 2]) << 16
	i |= uint64(buff[cursor + 3]) << 24
	i |= uint64(buff[cursor + 4]) << 32
	i |= uint64(buff[cursor + 5]) << 40
	i |= uint64(buff[cursor + 6]) << 48
	i |= uint64(buff[cursor + 7]) << 56
	return cursor + 8, i
}

func ReadLength(buff []byte, cursor int) (int, uint64) {
	length := buff[cursor]
	cursor++
	switch length {
	case 251:
		return cursor, 0
	case 252:
		cursor, u16 := ReadUB2(buff, cursor)
		return cursor, uint64(u16)
	case 253:
		cursor, u24 := ReadUB3(buff, cursor)
		return cursor, uint64(u24)
	case 254:
		cursor, u64 := ReadUB8(buff, cursor)
		return cursor, u64
	default:
		return cursor, uint64(length)

	}
}

func ReadString(buff []byte, cursor int) (int, string) {
	cursor, tmp := ReadBytes(buff, cursor, len(buff) - cursor)
	return cursor, string(tmp)
}

func ReadStringWithNull(buff []byte, cursor int) (int, string) {
	cursor, tmp := ReadWithNull(buff, cursor)
	return cursor, string(tmp)
}

func ReadLengthString(buff []byte, cursor int) (int, string) {
	cursor, strLen := ReadLength(buff, cursor)
	cursor, tmp := ReadBytes(buff, cursor, int(strLen))
	return cursor, string(tmp)
}

func ReadWithNull(buff []byte, cursor int) (int, []byte) {
	ret := []byte{}
	for ; ;  {
		if buff[cursor] != 0 {
			ret = append(ret, buff[cursor])
			cursor++
		} else {
			cursor++
			break
		}
	}
	return cursor, ret
}

func ReadBitSet(buff []byte, cursor int, length int, bigEndian bool) (int, bitset.BitSet) {
	cursor, bytes := ReadBytes(buff, cursor, (length + 7) >> 3)
	if bigEndian == false {
		bytes = ByteReverse(bytes)
	}
	var bs bitset.BitSet
	for i := 0; i < length; i++ {
		if (bytes[i >> 3] & (1 << (uint(i) % 8))) != 0 {
			bs.Set(uint(i))
		}
	}
	return cursor, bs
}
