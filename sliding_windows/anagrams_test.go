package sliding_windows

import "testing"

func TestSubstringAnagrams(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		output int
	}{
		{"caabab", "aba", 2},
	}
	for _, test := range tests {
		o := substringAnagram(test.s, test.t)
		if o != test.output {
			t.Errorf("substringAnagram(%v, %v) = %v, want %v", test.s, test.t, o, test.output)
		}
	}
}

func substringAnagram(s string, t string) interface{} {
	lenS, lenT := len(s), len(t)
	if lenS < lenT {
		return 0
	}
	count := 0
	expectedFreqs, windowFreqs := [26]int{}, [26]int{}

	// Populate 'expectedFreqs' with the characters in string 't'.
	for _, c := range t {
		expectedFreqs[int(c)-int('a')] += 1
	}

	left, right := 0, 0

	for right < lenS {

		// Add the character at the right pointer to 'windowFreqs'
		// before sliding window.
		windowFreqs[int(s[right])-int('a')] += 1

		if right-left+1 == lenT {
			if expectedFreqs == windowFreqs {
				count += 1
			}

			windowFreqs[int(s[left])-int('a')] -= 1
			left += 1
		}
		right += 1
	}
	return count
}
