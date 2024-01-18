package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	prices   []int
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, maxProfit(test.prices))
	}
}

func TestMore(t *testing.T) {
	tests := []Test{
		{
			prices:   []int{},
			expected: 0,
		},
		{
			prices:   []int{1, 1, 1, 1},
			expected: 0,
		},
		{
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			prices:   []int{9, 2, 8, 1, 9, 3, 4},
			expected: 8,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, maxProfit(test.prices))
	}
}
