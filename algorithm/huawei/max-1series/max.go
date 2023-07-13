package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	fmt.Println(longestOnes(nums, k))
}

func longestOnes(nums []int, k int) int {
	var result int
	for left, _ := range nums {
		cnt0 := 0
		right := left
		for right < len(nums) {
			if nums[right] == 0 {
				if cnt0 >= k {
					break
				}
				cnt0 += 1 - nums[right]
			}
			right++
		}
		result = max(result, right-left)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
