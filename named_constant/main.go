package main

import "fmt"

// 名前付き定数
func main() {
	// 型のある定数
	const n int = 100
	// 型のない定数
	const m = 100
	// 定数式の使用
	const s = "Hello, " + "世界"
	// まとめる
	const (
		x = 100
		y = 200
	)
	fmt.Println(n, m, s, x, y)

}
