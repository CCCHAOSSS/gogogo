package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap", cap(s))

	var t []string
	fmt.Println("emp:", t, "len:", len(t), "cap", cap(t))
}

func split(sum int) (int, int) {
	var x = sum * 4 / 9
	y := sum - x
	return x, y
}
