# 基于GO语言开发的mysql二进制协议 #

## 简介 ##
没啥，随便写写，目前仅仅是为了实现mysql协议而已，以后可能会基于此做mysql中间件

## 进度 ##
##### 2017.09.30 #####
* 完成util,common包的代码coding
* 完成HandsharkPacket解析

##### 2017.10.11 #####
* 完成AuthPacket解析
* 完成RegisterSlavePacket解析
* 完成ResultSetPacket、ResultHeader、Field解析
* 完成OKPacket解析
* 完成TABLE_MAP_EVENT解析

##### 下一步 #####
* 完成WRITE_ROWS_EVENT解析
* 完成UPDATE_ROWS_EVENT解析
* 完成DELETE_ROWS_EVENT解析

## 规划 ##
由于本人还在上班，总目标2017年底完成
* stone 1: 2017年11月，完成登录，主从复制相关协议开发
* stone 2: 2017年12月，完成mysql读写等相关协议开发

## Author ##
<ericniu>947847775@qq.com