package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// io.Reader/io.Writerでストリームを自由に操る

// io.TeeReader()は、読み込まれた内容を別のio.Writerに書き出す

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)
	// けどバッファに残ってる
	fmt.Println(buffer.String())
}
