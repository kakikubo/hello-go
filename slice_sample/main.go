package main

// 【TRY】スライスの利用
func main() {
	// n1 := 19
	// n2 := 86
	// n3 := 1
	// n4 := 12

	// sum := n1 + n2 + n3 + n4
	ns := []int{19, 86, 1, 12}
	sum := 0
	for _, n := range ns {
		sum += n
	}
	println(sum)
}
