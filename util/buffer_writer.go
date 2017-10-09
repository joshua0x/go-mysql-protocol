package util

func WriteByte(buf []byte, b byte) []byte {
	buf = append(buf, b)
	return buf
}

func WriteBytes(buf []byte, from []byte) []byte {
	for _, v := range from  {
		buf = append(buf, v)
	}
	return buf
}

func WriteUB2(buf []byte, i uint16) []byte {
	buf = append(buf, byte(i & 0xFF))
	buf = append(buf, byte((i >> 8) & 0xFF))
	return buf
}

func WriteUB3(buf []byte, i uint32) []byte {
	buf = append(buf, byte(i & 0xFF))
	buf = append(buf, byte((i >> 8) & 0xFF))
	buf = append(buf, byte((i >> 16) & 0xFF))
	return buf
}

func WriteUB4(buf []byte, i uint32) []byte {
	buf = append(buf, byte(i & 0xFF))
	buf = append(buf, byte((i >> 8) & 0xFF))
	buf = append(buf, byte((i >> 16) & 0xFF))
	buf = append(buf, byte((i >> 24) & 0xFF))
	return buf
}

func WriteUB8(buf []byte, i uint64) []byte {
	buf = append(buf, byte(i & 0xFF))
	buf = append(buf, byte((i >> 8) & 0xFF))
	buf = append(buf, byte((i >> 16) & 0xFF))
	buf = append(buf, byte((i >> 24) & 0xFF))
	buf = append(buf, byte((i >> 32) & 0xFF))
	buf = append(buf, byte((i >> 40) & 0xFF))
	buf = append(buf, byte((i >> 48) & 0xFF))
	buf = append(buf, byte((i >> 56) & 0xFF))
	return buf
}

func WriteLength(buf []byte, length int64) []byte {
	if length <= 251 {
		buf = WriteByte(buf, byte(length))
	} else if length < 0x10000 {
		buf = WriteByte(buf,252)
		buf = WriteUB2(buf, uint16(length))
	} else if length < 0x1000000 {
		buf = WriteByte(buf,253)
		buf = WriteUB3(buf, uint32(length))
	} else {
		buf = WriteByte(buf,254)
		buf = WriteUB8(buf, uint64(length))
	}
	return buf
}

func WriteWithNull(buf []byte, from []byte) []byte {
	buf = WriteBytes(buf, from)
	buf = append(buf, byte(0))
	return buf
}

func WriteWithLength(buf []byte, from []byte) []byte {
	length := len(from)
	buf = WriteLength(buf, int64(length))
	buf = WriteBytes(buf, from)
	return buf
}
