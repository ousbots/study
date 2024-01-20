package implement_queue_using_stacks

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type Test struct {
	commands []string
	args     []int
	expected []string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			commands: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
			args:     []int{-1, 1, 2, -1, -1, -1},
			expected: []string{"", "", "", "1", "1", "false"},
		},
		{
			commands: []string{"MyQueue", "empty", "push", "empty", "pop", "empty"},
			args:     []int{-1, -1, 1, -1, -1, -1},
			expected: []string{"", "true", "", "false", "1", "true"},
		},
		{
			commands: []string{"MyQueue", "push", "push", "push", "pop", "pop", "pop"},
			args:     []int{-1, 1, 2, 3, -1, -1, -1},
			expected: []string{"", "", "", "", "1", "2", "3"},
		},
		{
			commands: []string{"MyQueue", "push", "peek", "push", "peek", "push", "peek"},
			args:     []int{-1, 1, -1, 2, -1, 3, -1},
			expected: []string{"", "", "1", "", "1", "", "1"},
		},
		{
			commands: []string{"MyQueue", "push", "peek", "pop", "push", "peek", "pop", "push", "peek", "pop"},
			args:     []int{-1, 1, -1, -1, 2, -1, -1, 3, -1, -1},
			expected: []string{"", "", "1", "1", "", "2", "2", "", "3", "3"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, len(test.commands), len(test.args))

		var this *MyQueue
		var output []string

		for i := range test.commands {
			command := test.commands[i]
			arg := test.args[i]

			if command == "MyQueue" {
				temp := Constructor()
				this = &temp
				output = append(output, "")
			}

			if command == "push" {
				assert.NotNil(t, this, "can't push on nil")
				this.Push(arg)
				output = append(output, "")
			}

			if command == "pop" {
				assert.NotNil(t, this, "can't pop on nil")
				val := this.Pop()
				output = append(output, strconv.Itoa(val))
			}

			if command == "peek" {
				assert.NotNil(t, this, "can't peek on nil")
				val := this.Peek()
				output = append(output, strconv.Itoa(val))
			}

			if command == "empty" {
				assert.NotNil(t, this, "can't empty on nil")
				isEmpty := this.Empty()
				output = append(output, strconv.FormatBool(isEmpty))
			}
		}

		assert.Equal(t, test.expected, output, "test commands %v args %v", test.commands, test.args)
	}
}
