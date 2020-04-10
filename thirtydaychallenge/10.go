package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3292

type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, mins: []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 || x <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
