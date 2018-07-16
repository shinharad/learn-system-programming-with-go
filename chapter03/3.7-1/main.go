package main

import (
	"bytes"
	"io"
	"os"
)

// io.Reader/io.Writerでストリームを自由に操る

// io.MultiReaderは引数で渡されたio.Readerのすべての入力がつながっているかのように動作する

func main() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	// すべてのreaderをつなげた出力を表示
	io.Copy(os.Stdout, reader)
}
