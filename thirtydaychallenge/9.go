package main

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/529/week-2/

func parse(str string) string {
	var stack []rune
	for _, r := range str {
		if r == '#' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if r != '#' {
			stack = append(stack, r)
		}
	}
	parsedStr := ""
	for _, s := range stack {
		parsedStr += string(s)
	}
	return parsedStr
}

func backspaceCompare(S string, T string) bool {
	return parse(S) == parse(T)
}
