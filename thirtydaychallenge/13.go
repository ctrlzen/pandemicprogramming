package main

// https://leetcode.com/submissions/detail/324333389/?from=/explore/other/card/30-day-leetcoding-challenge/529/week-2/3298/

type arrayDetail struct {
	startIndex int
	maxLen     int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMaxLength(nums []int) int {

	maxLen := 0
	hashMap := map[int]*arrayDetail{}
	runningSums := []int{0}

	for i, num := range nums {
		if num == 0 {
			runningSums = append(runningSums, runningSums[i]-1)
		} else {
			runningSums = append(runningSums, runningSums[i]+1)
		}
	}

	for i, sum := range runningSums {
		if val, ok := hashMap[sum]; ok {
			val.maxLen = i - val.startIndex
		} else {
			hashMap[sum] = &arrayDetail{
				startIndex: i,
				maxLen:     0,
			}
		}
		maxLen = max(maxLen, hashMap[sum].maxLen)
	}

	return maxLen
}
