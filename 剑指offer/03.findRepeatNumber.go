package main

import (
	"fmt"
	"sort"
)

// 方法一
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		if _, ok := m[value]; ok {
			return value
		} else {
			m[value] = 1
		}
	}
	return -1
}

// 方法二
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i, numSize := 0, len(nums); i < numSize-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println(num)

	num2 := findRepeatNumber2(nums)
	fmt.Println(num2)
}
