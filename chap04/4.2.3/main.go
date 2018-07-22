package main

import (
	"fmt"
	"math"
)

// チャネルとfor文との連携

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
				if !found {
					result <- i
				}
			}
		}
		close(result)
	}()
	return result
}

func main() {
	// 「値がくるたびにforループが回る、個数が未定の動的配列」のように扱える
	// このforループはチャネルがオープンしている間は回り続けるが、チャネルがクローズされたら止まる
	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}
