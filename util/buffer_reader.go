package util

func ReadBytes(buff []byte, cursor int, offset int) (int, []byte) {
	return cursor + offset, buff[cursor:cursor + offset]
}

func ReadByte(buff []byte, cursor int) (int, byte) {
	return cursor + 1, buff[cursor]
}

func ReadUB2(buff []byte, cursor int) (int, uint16) {
	i := uint16(buff[cursor] & 0xFF)
	i |= uint16(buff[cursor + 1] & 0xFF) << 8
	return cursor + 2, uint16(i)
}

func ReadUB3(buff []byte, cursor int) (int, uint32) {
	i := uint32(buff[cursor] & 0xFF)
	i |= uint32((buff[cursor + 1] & 0xFF) << 8)
	i |= uint32((buff[cursor + 2] & 0xFF) << 16)
	return cursor + 3, i
}

func ReadUB4(buff []byte, cursor int) (int, uint32) {
	i := uint32(buff[cursor] & 0xFF)
	i |= uint32((buff[cursor + 1] & 0xFF) << 8)
	i |= uint32((buff[cursor + 2] & 0xFF) << 16)
	i |= uint32((buff[cursor + 3] & 0xFF) << 24)
	return cursor + 4, i
}

func ReadUB8(buff []byte, cursor int) (int, uint64) {
	i := uint64(buff[cursor] & 0xFF)
	i |= uint64((buff[cursor + 1] & 0xFF) << 8)
	i |= uint64((buff[cursor + 2] & 0xFF) << 16)
	i |= uint64((buff[cursor + 3] & 0xFF) << 24)
	i |= uint64((buff[cursor + 4] & 0xFF) << 32)
	i |= uint64((buff[cursor + 5] & 0xFF) << 40)
	i |= uint64((buff[cursor + 6] & 0xFF) << 48)
	i |= uint64((buff[cursor + 7] & 0xFF) << 56)
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
