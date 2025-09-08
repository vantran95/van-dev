package main

import "fmt"

func isValid(s string) bool {
	var slice []string

	for _, char := range s {
		if string(char) == "(" || string(char) == "{" || string(char) == "[" {
			slice = append(slice, string(char))
		} else {
			if len(slice) == 0 {
				return false
			}

			if string(char) == ")" {
				if string(slice[len(slice)-1]) != "(" {
					return false
				} else {
					slice = slice[:len(slice)-1]
				}
			} else if string(char) == "]" {
				if string(slice[len(slice)-1]) != "[" {
					return false
				} else {
					slice = slice[:len(slice)-1]
				}
			} else if string(char) == "}" {
				if string(slice[len(slice)-1]) != "{" {
					return false
				} else {
					slice = slice[:len(slice)-1]
				}
			}
		}
	}

	if len(slice) > 0 {
		return false
	}

	return true
}

func main() {
	s := "["

	fmt.Println(isValid(s))
}
