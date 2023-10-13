package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{111311, 1113}
	fmt.Println(largestNumber(nums))
}

type NumberArray []string

func largestNumber(nums []int) string {
	var strNums NumberArray
	for _, e := range nums {
		strNums = append(strNums, strconv.Itoa(e))
	}
	sort.Sort(strNums)
	result := strings.Join(strNums, "")
	if strings.HasPrefix(result, "0") {
		result = "0"
	}
	return result
}

func (na NumberArray) Len() int {
	return len(na)
}

func (na NumberArray) Less(i, j int) bool {
	return compare(na[i], na[j]) > 0
}

func (na NumberArray) Swap(i, j int) {
	tmp := na[i]
	na[i] = na[j]
	na[j] = tmp
}

func compare(str1, str2 string) int {
	con1, con2 := str1+str2, str2+str1
	return strings.Compare(con1, con2)
}
