package main

import "fmt"

var aa = ""

func hehe() {
	func() {
		fmt.Println(aa)
	}()
}

func main() {
	p := hehe
	aa = "aa"
	p()

}
