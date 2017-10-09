package util

import "fmt"

func WriteNoticeLog(log string) {
	fmt.Println("Notice: ", log)
}

func WriteErrorLog(log string) {
	fmt.Println("Error: ", log)
}