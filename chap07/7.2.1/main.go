package main

import (
	"fmt"
	"net"
)

// UDPとTCPの処理の流れの違い
// サーバー側の実装

func main() {
	fmt.Println("Server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 15000)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %d: %s\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
