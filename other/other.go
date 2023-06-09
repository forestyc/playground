package main

import (
	"fmt"
	"strings"
)

var aa = ""

func hehe() {
	func() {
		fmt.Println(aa)
	}()
}

func main() {
	str := "\n\n\n\naaaa"
	fmt.Println(strings.TrimPrefix(str, "\n"))
}
