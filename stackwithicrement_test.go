package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CustomStack struct {
	maxSize int
	Stack   []int
}

func Constructor2(maxSize int) CustomStack {
	return CustomStack{maxSize: maxSize}
}

func (this *CustomStack) Push(x int) {

	if len(this.Stack) < this.maxSize {
		this.Stack = append(this.Stack, x)
	}

}

func (this *CustomStack) Pop() int {

	if len(this.Stack) > 0 {
		popped := this.Stack[len(this.Stack)-1]
		this.Stack = this.Stack[:len(this.Stack)-1]
		return popped
	}

	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	var l int
	if len(this.Stack) < k {
		l = len(this.Stack)
	} else {
		l = k
	}

	for i := 0; i < l; i++ {
		this.Stack[i] += val
	}
}

func TestStack(t *testing.T) {
	stack := Constructor2(3)

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Increment(5, 100)
	stack.Increment(2, 100)

	assert.Equal(t, []int{201, 202, 103}, stack.Stack)
}
