package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "UC200923   "

	fmt.Println(regexp.MatchString("[a-zA-Z]+\\d{6}", str))
}
