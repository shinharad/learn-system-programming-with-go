package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// TCPソケットを使ったHTTPサーバー
//
// Go言語のソケットを使ってHTTP/1.0相当の送受信を実現する
// HTTP/1.0では、1セットの通信が終わるたびにTCPのコネクションが切れる仕様

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 非同期に実行される
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()

	}
}
