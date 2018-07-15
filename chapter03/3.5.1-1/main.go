package main

import (
	"io"
	"os"
	"strings"
)

// 必要な部位を切り出すio.LimitReader
func main() {
	reader := strings.NewReader("Readerの出力内容は文字列で渡す")
	// 先頭の9バイトしか読み込まないようにする
	lReader := io.LimitReader(reader, 9)
	io.Copy(os.Stdout, lReader)
}
