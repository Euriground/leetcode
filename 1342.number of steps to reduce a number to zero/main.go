package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
}

// TODO: implement. See README.md for the problem statement.
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// random change
func numberOfSteps(num int) int {
	count := 0

	for {
		if num == 0 {
			break
		}

		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}

		count++
	}
	return count
}
