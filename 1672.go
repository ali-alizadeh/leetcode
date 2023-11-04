// https://leetcode.com/problems/richest-customer-wealth

package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, customer := range accounts {
		wealth := 0
		for _, money := range customer {
			wealth += money
		}
		if wealth > max {
			max = wealth
		}
	}
	return max
}
