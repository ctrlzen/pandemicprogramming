package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/529/week-2/3293/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lHeight := height(node.Left)
	rHeight := height(node.Right)
	return max(lHeight, rHeight) + 1

}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := height(root.Left) + height(root.Right)
	ldiameter := diameterOfBinaryTree(root.Left)
	rdiameter := diameterOfBinaryTree(root.Right)

	return max(d, max(ldiameter, rdiameter))

}
