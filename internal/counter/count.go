package counter

import (
	"strings"
	"unicode/utf8"
)

func Count(content string) (byteCount, charCount, wordCount, lineCount int) {
	byteCount = len(content)
	charCount = utf8.RuneCountInString(content)

	if content == "" {
		return 0, 0, 0, 0
	}

	lines := strings.Split(content, "\n")
	lineCount = len(lines)
	if content[len(content)-1] == '\n' {
		lineCount--
	}

	for _, line := range lines {
		if line != "" {
			wordCount += len(strings.Fields(line))
		}
	}

	return
}
