package main

import "fmt"

func main() {
	t := Test{Id: 1}
	t.Echo()
	fmt.Println(t)
}

type Test struct {
	Id int
}

func (t *Test) Echo() {
	t.Id = 3
}
