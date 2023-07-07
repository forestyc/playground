package main

import (
	"fmt"
	"math/rand"
	"time"
)

var value []string = []string{
	"YES",
	"NO",
	"NA",
}

type Region struct {
	Grid    [][]string
	Remodel bool
}

func main() {
	//region := InitGrid(10, 10)
	//fmt.Println(region)
	//fmt.Println(Remodel(region))
	var row1, row2, row3, row4, row5, row6 string
	//fmt.Scanf("%s %s %s\n%s %s %s\n", &row1, &row2, &row3, &row4, &row5, &row6)
	var input string
	count, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input, count)
	fmt.Println(row1, row2, row3, row4, row5, row6)
}

func RodomValue() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := rand.Intn(3)
	return value[idx]
}

func InitGrid(row, column int) *Region {
	region := &Region{}
	region.Grid = make([][]string, row)
	for i := 0; i < row; i++ {
		region.Grid[i] = make([]string, column)
		for j := 0; j < column; j++ {
			region.Grid[i][j] = RodomValue()
			if region.Grid[i][j] == "YES" {
				region.Remodel = true
			}
		}
	}
	return region
}

func Remodel(region *Region) int {
	if region != nil || region.Remodel == false {
		return -1
	}
	var day int
	return day
}
