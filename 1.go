// https://leetcode.com/problems/two-sum

package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if index, ok := numsMap[target-num]; ok {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return []int{0, 0}
}

func twoSumWithRange(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
		if index, ok := numsMap[target-num]; ok {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return []int{}
}
