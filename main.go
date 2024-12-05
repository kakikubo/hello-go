package main

import (
	"fmt"
	"github.com/kakikubo/hello-go/greeting"
)

func main() {
	fmt.Println("Hello, World!")

	// バイト列と文字列を表すstring型とが次のように相互に変換可能
	byteArray := []byte("ASCII")
	fmt.Printf("%#v\n", byteArray) // []byte{0x41, 0x53, 0x43, 0x49, 0x49}
	str := string([]byte{0x41, 0x53, 0x43, 0x49, 0x49})
	fmt.Printf("%#v\n", str) // "ASCII"

	// greetingパッケージのDo関数を呼び出す
	message := greeting.Do()
	fmt.Println(message)
}
