package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3286/

func moveZeroes(nums []int) {

	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

}
