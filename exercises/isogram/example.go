package isogram

import "strings"

// IsIsogram example answer.
func IsIsogram(word string) bool {
	if word == "" {
		return false
	}
	lowerWord := strings.ToLower(word)
	for i, r := range lowerWord {
		if r != ' ' && r != '-' && i < strings.LastIndex(lowerWord, string(r)) {
			return false
		}
	}
	return true
}
