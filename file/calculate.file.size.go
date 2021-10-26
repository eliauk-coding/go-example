package main

import (
	"fmt"
	"os"
)

const (
	ConstGSize = 1024 * 1024 * 1024 // G
	ConstMSize = 1024 * 1024        // M
	ConstKSize = 1024               // K
)

func main() {
	fi, err := os.Stat("test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fileSize := fi.Size()
	// 计算G
	var sizeStr string
	if fileSize%ConstGSize != fileSize {
		sizeStr = fmt.Sprintf("%0.2fGB", float64(fileSize)/ConstGSize)
	} else if fileSize%ConstMSize != fileSize {
		sizeStr = fmt.Sprintf("%0.2fMB", float64(fileSize)/ConstMSize)
	} else if fileSize%ConstKSize != fileSize {
		sizeStr = fmt.Sprintf("%0.2fKB", float64(fileSize)/ConstKSize)
	} else {
		sizeStr = fmt.Sprintf("%dB", fileSize)
	}

	fmt.Println("file size is ", sizeStr, err)
}
