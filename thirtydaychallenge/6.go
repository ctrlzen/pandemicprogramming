package main

import "sort"

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3288/

func getAnagramKey(word string) string {
	chars := []rune{}
	for _, char := range word {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func groupAnagrams(words []string) [][]string {

	groups := [][]string{}
	anagrams := map[string][]string{}

	for _, word := range words {
		anagramKey := getAnagramKey(word)
		anagrams[anagramKey] = append(anagrams[anagramKey], word)
	}

	for _, anagramGroup := range anagrams {
		groups = append(groups, anagramGroup)
	}

	return groups

}
