package main

import (
	"net/http"
	"os"
)

// Request構造体
// 2.4.4ではHTTPプロトコルを手書きしたが、Request構造体を使えば、書き間違いなどのミスが減る
func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます ")
	request.Write(os.Stdout)
}
