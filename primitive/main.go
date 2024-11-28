package main

func main() {
	// 【TRY】組み込み型(数値)
	// https: //docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_111
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / float32(3)
	if avg > 4.5 {
		println("good")
	}
}
