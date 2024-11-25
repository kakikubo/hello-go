package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	num := 10
	str := "Hello, World!"
	fl := 3.14
	fmt.Fprintf(file, "number: %d\nstring: %s\nfloat: %f\n", num, str, fl)
}
