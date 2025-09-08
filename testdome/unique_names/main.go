package main

import (
	"fmt"
	"slices"
)

// using libraries
func uniqueNames(a, b []string) []string {
	var result []string

	for _, name := range a {
		if slices.IndexFunc(result, func(s string) bool { return name == s }) < 0 {
			result = append(result, name)
		}
	}

	for _, name := range b {
		if slices.IndexFunc(result, func(s string) bool { return name == s }) < 0 {
			result = append(result, name)
		}
	}

	return result
}

// using map data structure
func uniNames(a, b []string) []string {

	// The key always unique so make sure not duplicated name in map::key
	uniNames := make(map[string]bool)
	for _, name := range a {
		uniNames[name] = true
	}

	for _, name := range b {
		uniNames[name] = true
	}

	var result []string
	for name := range uniNames {
		result = append(result, name)
	}

	return result
}
func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))

	fmt.Println(uniNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
