package insert_interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	intervals [][]int
	insert    []int
	valid     [][]int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			intervals: [][]int{{1, 3}, {6, 9}},
			insert:    []int{4, 5},
			valid:     [][]int{{1, 3}, {4, 5}, {6, 9}},
		},
		{
			intervals: [][]int{{1, 3}, {6, 9}},
			insert:    []int{-1, 0},
			valid:     [][]int{{-1, 0}, {1, 3}, {6, 9}},
		},
		{
			intervals: [][]int{{1, 3}, {6, 9}},
			insert:    []int{10, 12},
			valid:     [][]int{{1, 3}, {6, 9}, {10, 12}},
		},
		{
			intervals: [][]int{{1, 3}, {6, 9}},
			insert:    []int{2, 5},
			valid:     [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			insert:    []int{4, 8},
			valid:     [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			insert:    []int{0, 6},
			valid:     [][]int{{0, 7}, {8, 10}, {12, 16}},
		},
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			insert:    []int{4, 17},
			valid:     [][]int{{1, 2}, {3, 17}},
		},
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			insert:    []int{0, 20},
			valid:     [][]int{{0, 20}},
		},
		{
			intervals: [][]int{},
			insert:    []int{1, 2},
			valid:     [][]int{{1, 2}},
		},
	}

	for _, test := range tests {
		var intervals [][]int
		for _, interval := range test.intervals {
			intervals = append(intervals, []int{interval[0], interval[1]})
		}

		assert.Equal(t, test.valid, insert(test.intervals, test.insert), "intervals: %v insert: %v", intervals, test.insert)
	}
}
