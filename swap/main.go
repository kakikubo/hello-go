package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

// 複数戻り値の利用
func swap(a, b int) (int, int) {
	return b, a
}
