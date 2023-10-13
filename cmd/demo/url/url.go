package main

import (
	"fmt"
	"net/url"
)

func main() {
	target := "etcd://baal1990.asia:23790/add.rpc"
	u, err := url.Parse(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*u)
}
