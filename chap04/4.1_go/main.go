package main

import (
	"fmt"
	"time"
)

// 新しく作られる goroutine が呼ぶ関数
func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	go sub()
	fmt.Println("started sub()")
	time.Sleep(2 * time.Second)
}
