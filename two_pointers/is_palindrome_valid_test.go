package two_pointers

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeValid(t *testing.T) {
	tcs := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a dog! a panic in a pagoda", true},
		{"adc123", false},
		{"byteetyb", true},
		{"racecar", true},
	}
	for _, c := range tcs {
		got := testIsPalindromeValid(c.input)
		if !assert.Equal(t, c.want, got) {
			t.Errorf("input: %s, got %v, want %v", c.input, got, c.want)
		}
	}
}

func testIsPalindromeValid(input string) bool {
	left, right := 0, len(input)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(input[left])) && !unicode.IsDigit(rune(input[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(input[right])) && !unicode.IsDigit(rune(input[right])) {
			right--
		}
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}
	return true
}
