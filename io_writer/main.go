package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// ファイル出力
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()

	// 画面出力
	os.Stdout.Write([]byte("os.Stdout example\n"))

	// バッファ
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
	// 以下でも実行可能。
	buffer.WriteString("bytes.Buffer example2\n")
	fmt.Println(buffer.String())
	// ただし、WriteStringはbytes.Bufferのメソッドであるため、
	// io.Writerインタフェースを満たすos.Stdoutやos.Fileでは利用できない
	// io.WriteString関数を使えばキャストは不要になる
	io.WriteString(&buffer, "bytes.Buffer example3\n")
	fmt.Println(buffer.String())
}
