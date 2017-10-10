package main

import (
	"go-mysql-protocol/socket"
	"go-mysql-protocol/protocol"
	"fmt"
	"go-mysql-protocol/util"
	"os"
	"strconv"
	"time"
)

func main() {
	var ret, tmp, body []byte
	var serverID uint64 = 0xFFFFFF
	//建立连接
	conn := socket.GetSocket("127.0.0.1", 3306)

	//读取握手信息
	_, _, body = socket.ReadPacket(conn)
	hs := protocol.DecodeHandshark(body)

	//发送登录包
	tmp = protocol.EncodeLogin(hs, "root", "MhxzKhl", "test")
	socket.WritePacket(conn, 1, tmp)

	//读取登录返回
	_, _, body = socket.ReadPacket(conn)
	fmt.Printf("LogingReturn: %s\n", string(body))
	if body[0] == 0xFF {
		util.WriteErrorLog("mysql login fail.")
		os.Exit(1)
	} else {
		util.WriteNoticeLog("mysql login success.")
	}

	//发送注册从库请求
	ret = protocol.EncodeRegisterSlave(conn, "127.0.0.1", 3306, "root", "MhxzKhl", serverID)
	socket.WritePacket(conn, 0, ret)
	//读取注册结果
	_, _, body = socket.ReadPacket(conn)
	fmt.Printf("RegisterSlaveReturn: %s\n", string(body))
	if ret[0] == 0xFF {
		util.WriteErrorLog("register slave fail.")
		os.Exit(1)
	} else {
		util.WriteNoticeLog("register slave success")
	}

	//发送命令，查看binlog信息
	tmp = protocol.EncodeQuery("show master status")
	socket.WritePacket(conn, 0, tmp)

	rs := protocol.DecodeResultSet(conn)
	binlogFileName := rs.Data[0].Data[0]
	binlogPosition := rs.Data[0].Data[1]

	p, _ := strconv.Atoi(binlogPosition)
	body = protocol.EncodeBinlogDump(serverID, uint64(p), binlogFileName)
	fmt.Println(string(body))
	socket.WritePacket(conn, 0, body)

	for ; ;  {
		_, _, body = socket.ReadPacket(conn)
		if len(body) <= 0 {
			time.Sleep(1000000000)
			fmt.Print(".")
			continue
		}
		if body[0] == 0xFF {
			fmt.Println(string(body))
			break
		} else {
			fmt.Println(string(body))
		}
	}
}
