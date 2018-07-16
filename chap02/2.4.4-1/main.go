package main

import (
	"io"
	"net"
	"os"
)

// インターネットアクセスの送信
func main() {

	// net.Dial()関数は net.Conn インターフェースを返す
	// net.Conn の実体は、net.TCPConn 構造体のポインタ
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")

	// サーバーから返ってきたレスポンスを画面に出力する
	io.Copy(os.Stdout, conn)

}
