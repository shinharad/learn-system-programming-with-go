package main

import (
	"io"
	"os"
)

// ファイル入力
func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
