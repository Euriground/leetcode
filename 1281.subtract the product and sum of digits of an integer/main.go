package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

// TODO: implement. See README.md for the problem statement.
// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
func subtractProductAndSum(n int) int {
	stringify := strconv.Itoa(n)
	nums := []int{}

	for i := 0; i < utf8.RuneCountInString(stringify); i++ {
		num, _ := strconv.Atoi(string(stringify[i]))
		nums = append(nums, num)
	}
	sum := 0
	product := 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		product *= nums[i]
	}
	return product - sum
}
