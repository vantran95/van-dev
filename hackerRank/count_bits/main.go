package main

import (
	"encoding/json"
	"fmt"
)

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
	s := fmt.Sprintf("%b", num)

	var count int32
	for _, text := range s {
		if text == '1' {
			count++
		}
	}
	return count
}

func main() {
	result := countBits(126)

	fmt.Printf("Result: %d\n", result)
}
