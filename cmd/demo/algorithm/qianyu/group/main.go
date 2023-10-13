package main

import (
	"fmt"
	"math/rand"
	"time"
)

var GroupCache [][]int

type NameInfo struct {
	Id   int    // 原组别
	Name string // 名字
}

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
	GroupCache = make([][]int, len(GroupList))
	fmt.Println(Grouping(GroupList))
}

// Grouping 分组
// 对原组别groupList重新分组,并返回新组别result,新组别人员数[2,3]
func Grouping(groupList [][]string) (result [][]string) {
	oldGroup := groupList
	var group []NameInfo
	for {
		oldGroup, group = Pick(oldGroup)
		if len(group) == 0 {
			break
		} else if len(group) == 1 {
			newGroups := GroupCache[group[0].Id]
			for i := 0; i < len(result); i++ {
				eq := false
				for _, e := range newGroups {
					if i == e {
						eq = true
						break
					}
				}
				if !eq {
					result[i] = append(result[i], group[0].Name)
				}
			}
		} else {
			var names []string
			for _, e := range group {
				names = append(names, e.Name)
				// 记录原组别与新组别对应关系
				GroupCache[e.Id] = append(GroupCache[e.Id], len(result))
			}
			result = append(result, names)
		}
	}
	return result
}

// Pick 选人
// 从组中随机选出2个人，2个人来自不同组别
func Pick(groupList [][]string) (oldGroup [][]string, names []NameInfo) {
	var x, y, lastY int
	for i := 0; i < 2; i++ {
		if len(groupList) == 0 {
			break
		}
		y = RandInt(len(groupList))
		if len(groupList) > 1 {
			y = Offset(lastY, y, len(groupList)-1)
		}
		lastY = y
		x = RandInt(len(groupList[y]))
		names = append(names, NameInfo{Id: y, Name: groupList[y][x]})
		newNames := groupList[y][0:x]
		newNames = append(newNames, groupList[y][x+1:]...)
		if len(newNames) == 0 {
			head := groupList[0:y]
			tail := groupList[y+1:]
			groupList = append(head, tail...)
		} else {
			groupList[y] = newNames
		}
	}
	return groupList, names
}

// RandInt 根据范围n获取随机数
func RandInt(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

// Offset 偏移
// 偏移y，使其与lastY不同
func Offset(lastY, y, l int) int {
	if y == lastY {
		if y == 0 {
			y++
		} else if y == l {
			y--
		} else {
			y++
		}
	}
	return y
}
