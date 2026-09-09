package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	s := fmt.Sprint(x)
	i := 0
	j := len(s) - 1
	fmt.Println(len(s))
	for {
		if i > j {
			break
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
