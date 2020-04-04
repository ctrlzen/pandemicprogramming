package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/528/week-1/3284/

func isHappy(n int) bool {
	visited := map[int]int{}

	for true {
		digits := getDigits(n)
		sum := getSum(digits)

		if sum == 1 {
			return true
		}
		if _, ok := visited[n]; !ok {
			visited[n] = sum
			n = sum
		} else {
			break
		}

	}
	return false

}

func getDigits(n int) []int {
	digits := []int{}
	for n != 0 {
		digit := n % 10
		digits = append(digits, digit)
		n = n / 10
	}
	return digits
}

func getSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num * num

	}
	return sum
}
