package main

import (
	"fmt"
	"github.com/forestyc/playground/pkg/http"
)

func main() {
	client := http.NewClient(false)
	defer client.Close()
	resp, err := client.Do(
		"get",
		"https://www.baidu.com/robots.txt",
		nil,
		nil,
	)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}
