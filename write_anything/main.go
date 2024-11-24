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
	number := 10
	string := "Hello, World!"
	float := 3.14
	fmt.Fprintf(file, "number: %d\nstring: %s\nfloat: %f\n", number, string, float)
}
