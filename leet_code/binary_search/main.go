package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return -1
}

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	result := search(nums, 9)

	nums2 := []int{-1, 0, 3, 5, 9, 12}
	rs := search(nums2, 2)

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Result: %d\n", rs)
}
