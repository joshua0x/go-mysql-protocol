package main

import (
	"go-mysql-protocol/socket"
	"go-mysql-protocol/protocol"
	"fmt"
	"go-mysql-protocol/util"
	"os"
)

func main() {
	var ret, tmp, body []byte
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
	ret = protocol.EncodeRegisterSlave(conn, "127.0.0.1", 3306, "root", "MhxzKhl")
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

	ret = protocol.EncodeBinlogDump()
	socket.WritePacket(conn, 0, ret)
	for ; ;  {
		_, _, ret = socket.ReadPacket(conn)
		if ret == nil {
			continue
		}

		if ret[0] == 0xFF {
			e := protocol.DecodeError(ret)
			fmt.Printf("%+v\n", e)
		}
	}

}
