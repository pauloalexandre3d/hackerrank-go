package stacks

import "testing"

func TestValidParenthesisExpression(t *testing.T) {
	tcs := []struct {
		input  string
		output bool
	}{
		{
			input:  "([]{})",
			output: true,
		},
		{
			input:  "([]{)}",
			output: false,
		},
	}

	for _, tc := range tcs {
		if validParenthesisExpression(tc.input) != tc.output {
			t.Errorf("Expected %v, got %v", tc.output, validParenthesisExpression(tc.input))
		}
	}
}

func validParenthesisExpression(input string) bool {
	parenthesisMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune
	for _, char := range input {
		if _, ok := parenthesisMap[char]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if parenthesisMap[stack[len(stack)-1]] == char {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
