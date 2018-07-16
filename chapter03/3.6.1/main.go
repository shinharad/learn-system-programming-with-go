package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 改行/単語で区切る

// io.Readerによる入力ではbufio.Readerを使う方が比較的シンプル
// bufio.Readerは読み込んだ文字列を戻すこともできるため、
// テキストの構文解析器を自前で作る際のベースにできる

var source = `1行め
2行め
3行め `

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}

	// bufio.Scannerの例
	// 分割文字が削除されてしまうことに注意
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
