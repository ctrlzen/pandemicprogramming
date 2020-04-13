package main

import "sort"

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/529/week-2/3297/

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		// sort descending
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
		if stones[0] == stones[1] {
			stones = stones[2:]
		} else {
			stones[1] = abs(stones[1] - stones[0])
			stones = stones[1:]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
