// https://leetcode.com/problems/fizz-buzz

package main

import "strconv"

func fizzBuzz(n int) []string {
	answer := make([]string, n)

	for i := 1; i <= n; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			answer[i-1] = "FizzBuzz"
		} else if divisibleBy3 {
			answer[i-1] = "Fizz"
		} else if divisibleBy5 {
			answer[i-1] = "Buzz"
		} else {
			answer[i-1] = strconv.Itoa(i)
		}
	}
	return answer
}
