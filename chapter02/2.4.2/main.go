package main

import (
	"os"
)

// 画面出力
func main() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
