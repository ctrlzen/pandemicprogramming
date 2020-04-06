package main

import (
	"bytes"
	"fmt"
	"sort"
)

func groupAnagrams(values []string) [][]string {
	anagramsByKey := map[string][]string{}

	for _, value := range values {
		key := computeKey(value)
		fmt.Println(key)
		anagramsByKey[key] = append(anagramsByKey[key], value)
	}
	fmt.Println(anagramsByKey)

	var groups [][]string

	// for _, group := range anagramsByKey {
	// 	groups = append(groups, group)
	// }

	return groups
}

func computeKey(value string) string {
	var characters []rune

	for _, char := range value {
		characters = append(characters, char)
	}

	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})

	var buf bytes.Buffer

	for _, char := range characters {
		buf.WriteRune(char)
	}

	return buf.String()
}

func main() {
	fmt.Println(groupAnagrams([]string{"cat", "tac", "ate", "eat"}))
}
