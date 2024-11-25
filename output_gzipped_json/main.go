package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	gzipper := gzip.NewWriter(w)
	// gzip圧縮を行いながら圧縮前の出力を標準出力に出力
	// io.MultiWriterを使う
	writer := io.MultiWriter(os.Stdout, gzipper)
	encoder := json.NewEncoder(writer)
	// JSONの文字列変換
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
	// os.Stdoutにjsonを出力する
	gzipper.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
