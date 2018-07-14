package main

import (
	"io"
	"os"
)

// io.Writerのデコレータ
// io.MultiWriter()
// 複数の io.Writer を受け取り、それらすべてに対して、書き込まれた内容を同時に書き込むデコレータ
func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	// fileとos.Stdoutに対する書き込み
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
