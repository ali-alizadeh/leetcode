// https://leetcode.com/problems/running-sum-of-1d-array/

package main

func runningSum(nums []int) []int {
	numsLength := len(nums)
	result := make([]int, numsLength)
	result[0] = nums[0]

	for index := 1; index < numsLength; index++ {
		result[index] = nums[index] + result[index-1]
	}

	return result
}
