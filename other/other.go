package main

import "fmt"

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
	var c [][]string
	for _, e := range GroupList {
		tmp := make([]string, len(e))
		copy(tmp, e)
		c = append(c, tmp)
	}
	//copy(c, GroupList)
	fmt.Println(GroupList)
	fmt.Println(c)

	s := make([]string, 0)
	fmt.Println(s[0])
}
