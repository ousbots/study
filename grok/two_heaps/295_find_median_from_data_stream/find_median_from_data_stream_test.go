package find_median_from_data_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	commands []string
	values   []int
	valid    []float64
}

func run(commands []string, values []int) []float64 {
	if len(commands) != len(values) {
		panic("commands and values are not the same length")
	}

	var system MedianFinder
	output := []float64{}

	for i := range commands {
		switch commands[i] {
		case "MedianFinder":
			system = Constructor()
			output = append(output, -1)

		case "addNum":
			system.AddNum(values[i])
			output = append(output, -1)

		case "findMedian":
			output = append(output, system.FindMedian())

		}
	}

	return output
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			commands: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			values:   []int{-1, 1, 2, -1, 3, -1},
			valid:    []float64{-1, -1, -1, 1.5, -1, 2.0},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, run(test.commands, test.values), "commands: %v values: %v", test.commands, test.values)
	}
}
