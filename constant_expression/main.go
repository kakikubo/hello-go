package main

import "fmt"

// 定数式
func main() {
	fmt.Printf("%#v\n", 100+200)
	fmt.Printf("%#v\n", 1<<2)
	fmt.Printf("%#v\n", "Hello, "+"世界")
	fmt.Printf("%#v\n", !(10 == 20))
}
