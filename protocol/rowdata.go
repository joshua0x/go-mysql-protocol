package protocol

import (
	"go-mysql-protocol/util"
)

type RowData struct {
	Data []string
}

func DecodeRowData(buf []byte) RowData {
	var c int = 0
	var s string
	var rowData RowData

	rowData.Data = make([]string, 0)
	for ; ;  {
		if c >= len(buf) {
			break
		}
		c, s = util.ReadLengthString(buf, c)
		rowData.Data = append(rowData.Data, s)
	}
	return rowData
}
