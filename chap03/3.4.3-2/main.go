package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

// ネットワーク通信の読み込み
// http.ReadResponse()でHTTPのレスポンスをパースする
func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示
	fmt.Println(res.Header)
	// ボディーを表示
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
