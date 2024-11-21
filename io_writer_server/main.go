package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// io.Writerインタフェースを満たすhttp.ResponseWriterに書き込む
	io.WriteString(w, "http.ResponseWriter example\n")
}

func main() {
	// httpサーバが起動し、リクエストがあるとhandler関数が呼ばれる
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
