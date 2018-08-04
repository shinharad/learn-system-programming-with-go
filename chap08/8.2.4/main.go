package main

import (
	"path/filepath"
	"os"
	"fmt"
	"net"
)

// データグラム型のUnixドメインソケット

// UDP相当の使い方ができるデータグラム型のUnixドメインソケットの実装

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsoket-server")
	// エラーチェックは削除（存在しなかったらしなかったで問題ないので不要）
	os.Remove(path)
	fmt.Println("Server is tunnning at " + path)
	conn, err := net.ListenPacket("unixgram", path)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n",
			remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"),
			remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}