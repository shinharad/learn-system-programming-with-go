package main

import (
	"bytes"
	"fmt"
	"io"
)

// 書かれた内容を記憶しておくバッファ
func main() {
	var buffer bytes.Buffer

	buffer.Write([]byte("bytes.Buffer example\n"))

	// このようにも書けるが、io.Writerのメソッドではないため、他の構造体では使えない
	buffer.WriteString("bytes.Buffer example\n")

	// 代わりに io.WriteString() 関数を使えばキャストは不要
	io.WriteString(&buffer, "bytes.Buffer example\n")

	fmt.Println(buffer.String())
}
