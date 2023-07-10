package main

import "fmt"

func main() {
	forceField := [][]int{
		{4, 4, 6},
		{7, 5, 3},
		{1, 6, 2},
		{5, 6, 3},
	}
	fmt.Println(fieldOfGreatestBlessing(forceField))
}

type Area struct {
	Top    float32
	Left   float32
	Right  float32
	Bottom float32
}

func fieldOfGreatestBlessing(forceField [][]int) int {
	var areaList []Area
	for _, e := range forceField {
		areaList = append(areaList, toArea(e))
	}
	return 0
}

func toArea(forceField []int) Area {
	return Area{
		Top:    float32(forceField[1]) + float32(forceField[2])/2,
		Left:   float32(forceField[0]) - float32(forceField[2])/2,
		Right:  float32(forceField[0]) + float32(forceField[2])/2,
		Bottom: float32(forceField[1]) - float32(forceField[2])/2,
	}
}
