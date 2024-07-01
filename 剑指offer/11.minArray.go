package main

// 153
func findMin(number []int) int {
	left, right := 0, len(number)-1
	for left < right {
		mid := left + (right-left)/2
		if number[mid] < number[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return number[left]
}

// 154 && 11
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[left]
}

// func main() {

// }
