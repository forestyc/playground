package main

import "fmt"

func main() {
	m := append([]int{}, 1, 2, 3)
	m1 := make([]int, 1)
	copy(m1, m[:1])
	fmt.Println(m)
	fmt.Println(m1)
}
