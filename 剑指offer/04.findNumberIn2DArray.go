package main

// 方法一
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] > target {
			break
		}
		if matrix[i][len(matrix[i])-1] < target {
			continue
		}
		col := binarySearch(matrix[i], target)
		if col != -1 {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 方法二
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if target > matrix[row][col] {
			row++
		} else if target < matrix[row][col] {
			col--
		} else {
			return true
		}
	}
	return false
}

// func main() {

// }
