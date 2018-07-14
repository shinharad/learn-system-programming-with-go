package main

import (
	"fmt"
)

// Talker ... インターフェースを定義
type Talker interface {
	Talk()
}

// Greeter ... 構造体を宣言
type Greeter struct {
	name string
}

// Talk ...
// 構造体は Talker インターフェースで定義されているメソッドを持っている
// 副作用のあるメソッドではレシーバーの型をポインタ型 (g *Greeter) にするが、無いのでポインタにしなくても良い
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {

	// インターフェースの型を持つ変数を宣言
	var talker Talker

	// インターフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{"wozozo"}
	talker.Talk()

}
