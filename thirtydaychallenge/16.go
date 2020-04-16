package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/530/week-3/3301/

func checkValidString(str string) bool {

	var leftCount int
	var rightCount int

	if len(str) == 0 || str == "*" {
		return true
	}

	if len(str) == 1 {
		return false
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ')' {
			leftCount--
		} else {
			leftCount++
		}
		if leftCount < 0 {
			return false
		}
	}

	if leftCount == 0 {
		return true
	}

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '(' {
			rightCount--
		} else {
			rightCount++
		}
		if rightCount < 0 {
			return false
		}
	}

	return true
}
