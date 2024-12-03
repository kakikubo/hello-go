package main

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}

// 複数戻り値の利用
func swap2(a, b *int) {
	*a, *b = *b, *a
}
