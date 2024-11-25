package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encodeing/json",
		"hello":   "world",
	})

	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-Test", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
