package main

// 方法1
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return nums[i]
		}
	}
	return nums[len(nums)-1] + 1
}

// 方法2
func missingNumber2(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] != mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {

}
