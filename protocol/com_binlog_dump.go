package protocol

func EncodeBinlogDump() []byte {
	var tmp int64
	buf := []byte{}

	//命令
	buf = append(buf, 0x12)

	//二进制日志数据的起始位置
	tmp = 0x00
	buf = append(buf, byte(tmp & 0xFF))
	buf = append(buf, byte((tmp >> 8) & 0xFF))
	buf = append(buf, byte((tmp >> 16) & 0xFF))
	buf = append(buf, byte((tmp >> 24) & 0xFF))

	//二进制日志数据标志位
	tmp = 0x00
	buf = append(buf, byte(tmp & 0xFF))
	buf = append(buf, byte((tmp >> 8) & 0xFF))
	buf = append(buf, byte((tmp >> 16) & 0xFF))
	buf = append(buf, byte((tmp >> 24) & 0xFF))

	//从服务器ID
	var sid uint64 = 0xFFFFFF00
	buf = append(buf, byte(sid & 0xFF))
	buf = append(buf, byte((sid >> 8) & 0xFF))
	buf = append(buf, byte((sid >> 16) & 0xFF))
	buf = append(buf, byte((sid >> 24) & 0xFF))

	return buf
}
