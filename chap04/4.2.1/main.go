package main

import "fmt"

// チャネルの使用方法

func main() {
	fmt.Println("start sub()")

	// 終了を受け取るためのチャネル
	done := make(chan bool)

	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()
	fmt.Println("started sub()")

	// 終了を待つ
	<-done

	fmt.Println("all tasks are finished")
}
