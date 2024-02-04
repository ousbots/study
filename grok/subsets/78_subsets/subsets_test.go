package subsets

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Test struct {
	input []int
	valid [][]int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{1, 2, 3},
			valid: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			input: []int{0},
			valid: [][]int{{}, {0}},
		},
		{
			input: []int{9, 0, 3, 5, 7},
			valid: [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}},
		},
	}

	for _, test := range tests {
		output := subsets(test.input)
		assert.Equal(t, len(test.valid), len(output))

		for i := range output {
			sort.Slice(test.valid[i], func(a, b int) bool { return test.valid[i][a] < test.valid[i][b] })
			sort.Slice(output[i], func(a, b int) bool { return output[i][a] < output[i][b] })
		}

		sort.Slice(test.valid, func(a, b int) bool {
			if len(test.valid[a]) == len(test.valid[b]) {
				for x := range test.valid[a] {
					if test.valid[a][x] == test.valid[b][x] {
						continue
					}

					return test.valid[a][x] < test.valid[b][x]
				}
			}

			return len(test.valid[a]) < len(test.valid[b])
		})

		sort.Slice(output, func(a, b int) bool {
			if len(output[a]) == len(output[b]) {
				for x := range output[a] {
					if output[a][x] == output[b][x] {
						continue
					}

					return output[a][x] < output[b][x]
				}
			}

			return len(output[a]) < len(output[b])
		})

		assert.Equal(t, test.valid, output, "input: %v", test.input)
	}
}
