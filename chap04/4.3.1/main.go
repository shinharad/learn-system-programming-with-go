package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// OSからのシグナルをチャネルで受け取る

func main() {
	// サイズが1より大きいチャネルを作成
	signals := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C) を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルがくるまで待つ
	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
