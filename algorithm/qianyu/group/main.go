package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	GroupList := [][]string{
		{"少华", "少平", "少军", "少安", "少康"},
		{"福军", "福堂", "福民", "福平", "福心"},
		{"小明", "小红", "小花", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
	}
	fmt.Println(Grouping(GroupList))
}

func Grouping(input [][]string) (result [][]string) {
	oldGroup := input
	var group []string
	for {
		oldGroup, group = Pick(oldGroup)
		if len(group) == 0 {
			break
		} else if len(group) == 1 {
			y := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(result))
			result[y] = append(result[y], group...)
		}
		result = append(result, group)
	}
	return result
}

func Pick(input [][]string) (oldGroup [][]string, names []string) {
	var x, y int
	for i := 0; i < 2; i++ {
		y = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(input))
		x = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(input[y]))
		if y == 0 {
			break
		}
		names = append(names, input[y][x])
		newNames := input[y][0:x]
		newNames = append(newNames, input[y][x+1:]...)
		if len(newNames) == 0 {
			head := input[0:y]
			tail := input[y+1:]
			input = append(head, tail...)
		} else {
			input[y] = newNames
		}
	}
	return input, names
}
