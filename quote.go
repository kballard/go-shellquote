package shellquote

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// Join quotes each argument and joins them with a space.
// If passed to /bin/sh, the resulting string will be split back into the
// original arguments.
func Join(args ...string) string {
	var buf bytes.Buffer
	for i, arg := range args {
		if i != 0 {
			buf.WriteByte(' ')
		}
		quote(arg, &buf)
	}
	return buf.String()
}

const (
	specialChars      = "\\'\"`${[|&;<>()*?!"
	extraSpecialChars = " \t\n"
	prefixChars       = "~"
)

func quote(word string, buf *bytes.Buffer) {
	// We want to try to produce a "nice" output. As such, we will
	// backslash-escape most characters, but if we encounter a space, or if we
	// encounter an extra-special char (which doesn't work with
	// backslash-escaping) we switch over to quoting the whole word. We do this
	// with a space because it's typically easier for people to read multi-word
	// arguments when quoted with a space rather than with ugly backslashes
	// everywhere.
	if len(word) == 0 {
		// oops, no content
		buf.WriteString("''")
		return
	}

	if strings.ContainsAny(word, extraSpecialChars) {
		// Use single-quoting for nicer output
		// (Everything is okay except ' itself, which must be rewritten as '\'',
		// unless it's at the beginning or end, in which case it can be
		// simplified to \\')
		leftTrimmed := strings.TrimLeft(word, "'")
		numStartingQuotes := len(word) - len(leftTrimmed)
		trimmed := strings.TrimRight(leftTrimmed, "'")
		numEndingQuotes := len(leftTrimmed) - len(trimmed)

		for i := 0; i < numStartingQuotes; i++ {
			buf.WriteString("\\'")
		}
		buf.WriteByte('\'')
		for _, runeValue := range trimmed {
			if runeValue == '\'' {
				buf.WriteString("'\\''")
			} else {
				buf.WriteRune(runeValue)
			}
		}
		buf.WriteByte('\'')
		for i := 0; i < numEndingQuotes; i++ {
			buf.WriteString("\\'")
		}
	} else {
		// Use backslash escaping
		firstRune, remainderIndex := utf8.DecodeRuneInString(word)
		if strings.ContainsRune(specialChars, firstRune) || strings.ContainsRune(prefixChars, firstRune) {
			buf.WriteByte('\\')
		}
		buf.WriteRune(firstRune)
		for _, runeValue := range word[remainderIndex:] {
			if strings.ContainsRune(specialChars, runeValue) {
				buf.WriteByte('\\')
			}
			buf.WriteRune(runeValue)
		}
	}
}
