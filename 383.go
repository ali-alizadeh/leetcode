// https://leetcode.com/problems/ransom-note

package main

func canConstruct(ransomNote string, magazine string) bool {
	lettersCount := make(map[rune]int)

	for _, char := range magazine {
		if _, ok := lettersCount[char]; ok {
			lettersCount[char] += 1
		} else {
			lettersCount[char] = 1
		}
	}

	for _, char := range ransomNote {
		if value, ok := lettersCount[char]; ok && value != 0 {
			lettersCount[char] -= 1
		} else {
			return false
		}
	}

	return true
}
