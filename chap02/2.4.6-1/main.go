package main

import (
	"fmt"
	"os"
	"time"
)

// フォーマットしてデータをio.Writerに書き出す
func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
