package sliding_windows

import "testing"

func TestLongestSubstringWithUniqueCharacters(t *testing.T) {
	in := "abcba"
	out := 3

	got := longestSubstringWithUniqueCharacters(in)
	if got != out {
		t.Errorf("got %d, want %d", got, out)
	}

	got = longestSubstringWithUniqueCharactersOptimized(in)
	if got != out {
		t.Errorf("got %d, want %d", got, out)
	}
}

func longestSubstringWithUniqueCharacters(in string) int {
	maxLen := 0
	hashSet := make(map[byte]struct{})
	left, right := 0, 0

	for right < len(in) {

		for {
			if _, exists := hashSet[in[right]]; !exists {
				break
			}
			delete(hashSet, in[left])
			left += 1
		}

		hashSet[in[right]] = struct{}{}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		// Expand the window
		right += 1
	}
	return maxLen
}

func longestSubstringWithUniqueCharactersOptimized(in string) int {
	maxLen := 0
	prevIndexes := make(map[byte]int)
	left, right := 0, 0

	for right < len(in) {
		if prevIndex, exists := prevIndexes[in[right]]; exists {
			left = max(left, prevIndex+1)
		}
		prevIndexes[in[right]] = right
		maxLen = max(maxLen, right-left+1)
		right += 1
	}
	return maxLen
}
