package main

import (
	"fmt"
	"strings"
)

// データ型を指定して解析

// io.Readerから読み込んだデータは、今のところ単なるバイト列か文字列としてしか扱っていない
// io.Readerのデータを整数や浮動小数点数に変換するには、fmt.Fscan()を使う

var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
