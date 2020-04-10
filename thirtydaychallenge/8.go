package main

import "math"

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3290/

func getMiddleIdx(size float64) float64 {
	if int(size)%2 == 0 {
		return size / 2
	} else {
		return math.Ceil(size / 2)
	}
}

func getSize(head *ListNode) float64 {
	var size float64
	node := head
	for node.Next != nil {
		size++
		node = node.Next
	}
	return size
}

func middleNode(head *ListNode) *ListNode {

	size := getSize(head)
	middleIdx := getMiddleIdx(size)

	node := head
	var idx float64
	for node.Next != nil {
		if idx == middleIdx {
			return node
		}
		idx++
		node = node.Next
	}
	return node

}
