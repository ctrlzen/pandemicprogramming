package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/528/week-1/3285/

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	maxSum := nums[0]
	curSum := maxSum
	for i := 1; i < len(nums); i++ {
		curSum = Max(nums[i], curSum+nums[i])
		maxSum = Max(curSum, maxSum)
	}
	return maxSum
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
