package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/530/week-3/3303/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	output := make([][]int, m)

	/**
	 * each element value of the first row and column in output
	 * is the sum of the value of the corresponding element in
	 * grid and the previous element of output. we pre-assign
	 * the first element of output with the value of the first
	 * element in grid since the cost to reach it is zero
	 */

	// first element
	output[0] = append(output[0], grid[0][0])

	// first row
	for i := 1; i < n; i++ {
		output[0] = append(output[0], grid[0][i]+output[0][i-1])
	}

	// first column
	for i := 1; i < m; i++ {
		output[i] = append(output[i], grid[i][0]+output[i-1][0])
	}

	// fill rest of matrix
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			//min path sum formula
			output[i] = append(output[i], grid[i][j]+min(output[i-1][j], output[i][j-1]))
		}
	}
	// return last element
	return output[m-1][n-1]

}
