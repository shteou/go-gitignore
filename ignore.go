package ignore

import (
	"strings"
)

type Ignore struct {
	IgnoreEntries []IgnoreEntry
}

type IgnoreEntry struct {
	Kind     string
	Value    string
	Original string
}

// TrimRightSpace trims any trailing spaces unless preceeded by a \
func TrimRightSpace(s string) string {
	stop := len(s) - 1
	for ; stop >= 0 && s[stop] == ' '; stop-- {
	}

	if stop >= 0 && s[stop] == '\\' {
		stop++
	}

	return s[:stop+1]
}

// Unescape Unuscape any backslashes
func Unescape(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] != '\\' {
			res = res + string(s[i])
		} else if i < len(s)-1 && s[i+1] == '\\' {
			res = res + string(s[i])
			i++
		}
	}

	return res
}

// ParseIgnoreLines Parse the supplied lines into a Go representation of that those
// gitignore entries
func ParseIgnoreLines(lines []string) (*Ignore, error) {
	ignoreEntries := []IgnoreEntry{}

	for _, l := range lines {
		stripped := TrimRightSpace(l)

		if stripped == "" {
			ignoreEntries = append(ignoreEntries, IgnoreEntry{"Empty", l, l})
		} else if strings.HasPrefix(l, "#") {
			ignoreEntries = append(ignoreEntries, IgnoreEntry{"Comment", strings.TrimSpace(stripped[1:]), l})
		} else if strings.HasPrefix(l, "!") {
			ignoreEntries = append(ignoreEntries, IgnoreEntry{"NegatedPath", Unescape(stripped[1:]), l})
		} else {
			ignoreEntries = append(ignoreEntries, IgnoreEntry{"Path", Unescape(stripped), l})
		}
	}

	return &Ignore{ignoreEntries}, nil
}

// ParseIgnoreBytes Parse the supplied byte array into a Go representation of those
// gitignore entries
func ParseIgnoreBytes(bytes []byte) (*Ignore, error) {
	lines := strings.Split(string(bytes), "\n")
	return ParseIgnoreLines(lines)
}
