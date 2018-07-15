package main

import (
	"io"
	"os"
	"strings"
)

// 必要な部位を切り出すio.SectionReader
// 文字列からSectionの部分だけを切り出したReaderをまず作成し、
// それをすべてos.Stdoutに書き出している
// （実際には文字列を分けるためにio.SectionReaderを使うことはまずありません）
func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	SectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, SectionReader)
}
