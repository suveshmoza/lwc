package tests

import (
	"lwc/internal/counter"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			bytes int
			chars int
			words int
			lines int
		}
	}{
		{"Hello world", struct{ bytes, chars, words, lines int }{11, 11, 2, 1}},
		{"    ", struct{ bytes, chars, words, lines int }{4, 4, 0, 1}},
		{"", struct{ bytes, chars, words, lines int }{0, 0, 0, 0}},
		{"\n", struct{ bytes, chars, words, lines int }{1, 1, 0, 1}},
		{"Hello\nWorld", struct{ bytes, chars, words, lines int }{11, 11, 2, 2}},
		{"Hello\n\nWorld", struct{ bytes, chars, words, lines int }{12, 12, 2, 3}},
		{"word", struct{ bytes, chars, words, lines int }{4, 4, 1, 1}},
		{"Hello 😊", struct{ bytes, chars, words, lines int }{10, 7, 2, 1}},
	}

	for _, test := range tests {
		bytes, chars, words, lines := counter.Count(test.input)
		if bytes != test.expected.bytes {
			t.Errorf("For input %q: Expected bytes %d, got %d", test.input, test.expected.bytes, bytes)
		}
		if chars != test.expected.chars {
			t.Errorf("For input %q: Expected chars %d, got %d", test.input, test.expected.chars, chars)
		}
		if words != test.expected.words {
			t.Errorf("For input %q: Expected words %d, got %d", test.input, test.expected.words, words)
		}
		if lines != test.expected.lines {
			t.Errorf("For input %q: Expected lines %d, got %d", test.input, test.expected.lines, lines)
		}
	}
}
