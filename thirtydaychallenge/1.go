package main

import "sort"

//https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/528/week-1/3283/

func singleNumber(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l; i += 2 {
		if (i == l-1) || (nums[i] != nums[i+1]) {
			return nums[i]
		}

	}
	return 0

}
