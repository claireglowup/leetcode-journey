package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalendar(t *testing.T) {

	calendar := Constructor()
	assert.Equal(t, true, calendar.Book(10, 20))
	assert.Equal(t, false, calendar.Book(15, 25))
	assert.Equal(t, true, calendar.Book(20, 30))

}

type MyCalendar struct {
	Intervals [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, interval := range this.Intervals {
		if start < interval[1] && interval[0] < end {
			return false
		}
	}

	this.Intervals = append(this.Intervals, []int{start, end})

	return true
}
