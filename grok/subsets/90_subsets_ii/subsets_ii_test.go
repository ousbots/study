package subsets_ii

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
			input: []int{1, 2, 2},
			valid: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			input: []int{0},
			valid: [][]int{{}, {0}},
		},
	}

	for _, test := range tests {
		output := subsetsWithDup(test.input)

		sort.Slice(test.valid, func(a, b int) bool {
			if len(test.valid[a]) == len(test.valid[b]) {
				for i := range test.valid[a] {
					if test.valid[a][i] == test.valid[b][i] {
						continue
					}

					return test.valid[a][i] < test.valid[b][i]
				}

				return false
			}

			return len(test.valid[a]) < len(test.valid[b])
		})

		sort.Slice(output, func(a, b int) bool {
			if len(output[a]) == len(output[b]) {
				for i := range output[a] {
					if output[a][i] == output[b][i] {
						continue
					}

					return output[a][i] < output[b][i]
				}

				return false
			}

			return len(output[a]) < len(output[b])
		})

		assert.Equal(t, test.valid, output, "input: %v", test.input)
	}
}
