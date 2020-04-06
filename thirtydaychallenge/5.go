package main

//https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/528/week-1/3287/

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {

	profit := 0
	for i := 1; i < len(prices); i++ {
		profit += Max(prices[i]-prices[i-1], 0)
	}
	return profit

}
