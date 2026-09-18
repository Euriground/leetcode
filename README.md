# leetcode

My solutions to LeetCode problems, written in Go. Each problem lives in its own
folder named `<id>.<title>` containing a `README.md` (problem statement + link)
and a `main.go` with the solution.

## Problems

| # | Problem | Difficulty | Folder |
| --- | --- | --- | --- |
| 9 | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | Easy | [`9.palindrome number`](./9.palindrome%20number) |
| 66 | [Plus One](https://leetcode.com/problems/plus-one/) | Easy | [`66.plus one`](./66.plus%20one) |
| 202 | [Happy Number](https://leetcode.com/problems/happy-number/) | Easy | [`202.happy number`](./202.happy%20number) |
| 231 | [Power of Two](https://leetcode.com/problems/power-of-two/) | Easy | [`231.power of two`](./231.power%20of%20two) |
| 258 | [Add Digits](https://leetcode.com/problems/add-digits/) | Easy | [`258.add digits`](./258.add%20digits) |
| 1281 | [Subtract the Product and Sum of Digits of an Integer](https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/) | Easy | [`1281.subtract the product and sum of digits of an integer`](./1281.subtract%20the%20product%20and%20sum%20of%20digits%20of%20an%20integer) |
| 1342 | [Number of Steps to Reduce a Number to Zero](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/) | Easy | [`1342.number of steps to reduce a number to zero`](./1342.number%20of%20steps%20to%20reduce%20a%20number%20to%20zero) |

## Running a solution

Each solution is a standalone `package main`. From the repo root:

```bash
go run "./9.palindrome number/main.go"
```

The `.cursor/install.sh` bootstrap compiles every solution's `main.go` to catch
build errors.
