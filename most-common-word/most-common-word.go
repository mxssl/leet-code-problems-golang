package mostcommonword

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	// Convert paragraph to lower case
	paragraph = strings.ToLower(paragraph)

	// Convert paragraph to slice of strings
	paragraphSlice := strings.FieldsFunc(paragraph, splitHelper)

	// Count how often words occur in a paragraph
	paragraphWordFrequency := make(map[string]int)
	for _, word := range paragraphSlice {
		if frequency, found := paragraphWordFrequency[word]; found {
			paragraphWordFrequency[word] = frequency + 1
		} else {
			paragraphWordFrequency[word] = 1
		}
	}

	// Create a map with banned words
	bannedMap := make(map[string]bool)
	for _, b := range banned {
		bannedMap[b] = true
	}

	// Remove banned words from a result
	var maxFrequency int
	var result string
	for word, frequency := range paragraphWordFrequency {
		if !bannedMap[word] {
			frequency++
			if frequency > maxFrequency {
				maxFrequency = frequency
				result = word
			}
		}
	}

	return result
}

// Helper func for convert paragraph to slice of strings
func splitHelper(r rune) bool {
	if r == ' ' || r == '?' || r == '!' || r == '.' || r == ';' || r == ',' || r == '*' || r == '\'' {
		return true
	}
	return false
}
