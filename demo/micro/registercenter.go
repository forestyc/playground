package main

import (
	"context"
	"fmt"
	"github.com/forestyc/playground/pkg/distributed/registercenter"
	"github.com/forestyc/playground/pkg/distributed/registercenter/etcd"
	"strconv"
	"time"
)

var (
	pubs []registercenter.Publisher
	sub  registercenter.Subscriber
)

func main() {
	var pubCancels []context.CancelFunc
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		go G1(ctx, i)
		pubCancels = append(pubCancels, cancel)
	}
	time.Sleep(time.Second)
	ctx, subCancel := context.WithCancel(context.Background())
	go G2(ctx)
	time.Sleep(time.Second)
	fmt.Println("+++++++++++++++++++++++++init++++++++++++++++++++++++")
	for i := 0; i < 9; i++ {
		fmt.Println(sub.GetService("G1"))
	}
	fmt.Println("+++++++++++++++++++++++++init++++++++++++++++++++++++")

	fmt.Println("+++++++++++++++++++++++++del 2++++++++++++++++++++++++")
	pubCancels[2]()
	time.Sleep(5 * time.Second)
	for i := 0; i < 9; i++ {
		fmt.Println(sub.GetService("G1"))
	}
	fmt.Println("+++++++++++++++++++++++++del 2++++++++++++++++++++++++")
	time.Sleep(time.Hour)
	subCancel()
}

func G1(ctx context.Context, i int) {
	pub := &etcd.Publisher{}
	pub.Register(
		[]string{
			"81.70.188.168:12379",
			"140.143.163.171:12379",
			"101.42.23.168:12379",
		},
		registercenter.Service{
			Name: "G1." + strconv.Itoa(i),
			Url:  "http://g1/" + strconv.Itoa(i),
		}, 3)
	pubs = append(pubs, pub)
	for {
		select {
		case <-ctx.Done():
			pub.Close()
			fmt.Println("G1" + strconv.Itoa(i) + "exit")
			return
		}
	}
}

func G2(ctx context.Context) {
	sub = &etcd.Subscriber{}
	sub.Listen([]string{
		"81.70.188.168:12379",
		"140.143.163.171:12379",
		"101.42.23.168:12379",
	}, "G1")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("G2 exit")
			return
		}
	}
	sub.Close()
}
