package main

func main() {
	var a, b, c bool
	/*
		【TRY】組み込み型(真偽値)
			a | b | c | 結果
			--|---|---|-----
			0 | 0 | 0 | true
			1 | 0 | 0 | true
			1 | 1 | 0 | true
			0 | 1 | 0 | true
			0 | 0 | 1 | false
			1 | 0 | 1 | false
			1 | 1 | 1 | true
			0 | 1 | 1 | false
	*/
	a, b, c = false, true, true
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
