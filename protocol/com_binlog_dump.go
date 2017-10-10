package protocol

import "go-mysql-protocol/util"

func EncodeBinlogDump(serverID uint64, binlogPosition uint64, binlogFileName string) []byte {
	var tmp int64
	buf := []byte{}

	//命令
	buf = append(buf, 0x12)

	//二进制日志数据的起始位置
	buf = append(buf, byte(binlogPosition & 0xFF))
	buf = append(buf, byte((binlogPosition >> 8) & 0xFF))
	buf = append(buf, byte((binlogPosition >> 16) & 0xFF))
	buf = append(buf, byte((binlogPosition >> 24) & 0xFF))

	//二进制日志数据标志位
	tmp = 0x00
	buf = append(buf, byte(tmp & 0xFF))
	buf = append(buf, byte((tmp >> 8) & 0xFF))


	//从服务器ID
	buf = append(buf, byte(serverID & 0xFF))
	buf = append(buf, byte((serverID >> 8) & 0xFF))
	buf = append(buf, byte((serverID >> 16) & 0xFF))
	buf = append(buf, byte((serverID >> 24) & 0xFF))

	buf = util.WriteBytes(buf, []byte(binlogFileName))

	return buf
}
