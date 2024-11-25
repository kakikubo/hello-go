package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	records := [][]string{
		{"名前", "年齢", "出身地", "性別"},
		{"山本", "24", "兵庫", "男"},
		{"鈴木", "29", "神奈川", "女"},
		{"佐藤", "27", "鹿児島", "男"},
	}
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Println(err)
		}
	}
	w.Flush()
}
