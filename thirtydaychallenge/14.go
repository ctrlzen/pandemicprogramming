package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/529/week-2/3299/

func stringShift(s string, shifts [][]int) string {
	for _, shift := range shifts {
		switch direction := shift[0]; direction {
		case 0:
			chars := s[:shift[1]]
			s = s[shift[1]:] + chars
		case 1:
			chars := s[len(s)-shift[1]:]
			s = chars + s[:len(s)-shift[1]]
		}
	}
	return s
}
