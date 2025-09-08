package main

import "fmt"

func findMaxSum(numbers []int) int {
	// Validate if numbers list have less than 2 elements
	if len(numbers) < 2 {
		return 0
	}

	// set 2 elemenet largest
	max1, max2 := 0, 0

	// find 2 elements largest numbers in list
	for _, num := range numbers {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	return max1 + max2
}

func main() {
	// list := []int{5, 9, 7, 11} // Expect 20
	// list := []int{3, 9, 5, 2} // Expect 14
	// list := []int{8, 1, 3, 7} // Expect 15

	list := []int{12, 20, 99, 1} // Expect 119

	fmt.Println(findMaxSum(list))
}
