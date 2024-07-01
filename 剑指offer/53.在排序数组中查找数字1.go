package main

// 方法1
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

// 方法2
// 二分查找：先找第一次出现的位置，再找最后一次出现的位子， 二者之差再加1就是目标元素出现的次数
func binary_search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	firstPosition := findFirstPosition(nums, target)
	if firstPosition == -1 {
		return 0
	}

	lastPostion := findLastPosition(nums, target)
	return lastPostion - firstPosition + 1
}

func findFirstPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // [mid+1, right]
		} else {
			right = mid // [left, mid]
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findLastPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)/2 // +1向上取整
		if nums[mid] <= target {
			left = mid // [mid, right]
		} else {
			right = mid - 1 // [left, mid-1]
		}
	}
	return left
}

// func main() {

// }
