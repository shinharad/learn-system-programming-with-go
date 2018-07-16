package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// エンディアン変換
//
// バイナリ解析ではエンディアン変換が必要
// 現在主流のCPUはリトルエンディアン
// リトルエンディアンでは、10000という数値（16進表示で0x2710） があったときに、
// 小さい桁からメモリに格納される
// しかし、ネットワーク上で転送されるデータの多くは、大きい桁からメモリに格納されるビッグエンディアン
// そのため、多くの環境では、ネットワークで受け取ったデータをリトルエンディアンに修正する必要がある
func main() {
	// 32ビットのビッグエンディアンのデータ
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
