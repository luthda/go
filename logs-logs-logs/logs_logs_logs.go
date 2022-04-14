package logs

import (
	"strings"
	"unicode/utf8"
)

func Application(log string) string {
	mapOfRunes := map[rune]string {
			0x2757:  "recommendation",
			0x1f50d: "search",
			0x2600:  "weather",
	}

	for _, char := range log {
		if value, exists := mapOfRunes[char]; exists{
			return value
		}
	}

	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
