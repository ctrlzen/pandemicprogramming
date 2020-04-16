package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/530/week-3/3300/

func productExceptSelf(nums []int) []int {

	output := make([]int, len(nums))
	var product int

	// get all products to the left of each index
	product = 1
	for i := 0; i < len(nums); i++ {
		output[i] = product
		product *= nums[i]

	}

	// get all products to the right of each index
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = product * output[i]
		product *= nums[i]

	}

	return output
}
