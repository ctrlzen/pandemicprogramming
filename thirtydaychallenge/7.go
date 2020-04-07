package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/528/week-1/3289/

func countElements(arr []int) int {
	count := 0
	elements := map[int]int{}
	for _, num := range arr {
		elements[num]++
	}
	for _, num := range arr {
		if _, ok := elements[num+1]; ok {
			count++
		}
	}
	return count
}
