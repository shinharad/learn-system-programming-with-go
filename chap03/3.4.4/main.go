package main

import (
	"bytes"
	"strings"
)

// メモリに蓄えた内容をio.Readerとして読み出すバッファ
// 初期化の方法は何通りかあって、それぞれ初期データが必要かどうかや、初期化データの型に応じて使い分ける
func main() {

	// 空のバッファ
	var buffer1 bytes.Buffer
	// バイト列で初期化
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	// 文字列で初期化
	buffer3 := bytes.NewBufferString("初期文字列")

	_, _, _ = buffer1, buffer2, buffer3

	// bytes.Readerはbytes.NewReaderで作成
	bReader1 := bytes.NewReader([]byte{0x10, 0x20, 0x30})
	bReader2 := bytes.NewReader([]byte("文字列をバイト配列にキャストして設定"))

	// strings.Readerはstrings.NewReader()関数で作成
	sReader := strings.NewReader("Readerの出力内容は文字列で渡す")

	_, _, _ = bReader1, bReader2, sReader

}
