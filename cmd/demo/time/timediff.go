package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var start, end string
	flag.StringVar(&start, "start", "", "start-time")
	flag.StringVar(&end, "end", "", "end-time")
	flag.Parse()

	startTime, _ := time.ParseInLocation("15:04", start, time.Local)
	endTime, _ := time.ParseInLocation("15:04", end, time.Local)

	fmt.Printf("result=[%.1f]\n", endTime.Sub(startTime).Hours()-1.5)
}
