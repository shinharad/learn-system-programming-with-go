package main

import (
	"os"
)

// ファイル出力
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}
