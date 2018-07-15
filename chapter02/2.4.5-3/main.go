package main

import (
	"bufio"
	"os"
)

// bufio.Writer
// 出力結果を一時的にためておいて、ある程度の分量ごとにまとめて書き出すbufio.Writerという構造体
// Flush()メソッドを呼ぶと、皇族のio.Writerに書き出す
// Flush()メソッドを呼ばないと、書き込まれたデータを腹に抱えたまま消滅するので要注意
func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
