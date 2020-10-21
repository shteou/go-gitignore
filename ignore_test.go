package ignore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreParseLine(t *testing.T) {
	result, _ := ParseIgnore([]string{".idea"})

	assert.Equal(t, len(result.IgnoreEntries), 1)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "Path")
	assert.Equal(t, result.IgnoreEntries[0].Original, ".idea")
	assert.Equal(t, result.IgnoreEntries[0].Value, ".idea")
}

func TestIgnoreParseNegation(t *testing.T) {
	result, _ := ParseIgnore([]string{"!.idea"})

	assert.Equal(t, len(result.IgnoreEntries), 1)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "NegatedPath")
	assert.Equal(t, result.IgnoreEntries[0].Original, "!.idea")
	assert.Equal(t, result.IgnoreEntries[0].Value, ".idea")
}

func TestIgnoreParseComment(t *testing.T) {
	result, _ := ParseIgnore([]string{"# Amazing"})

	assert.Equal(t, len(result.IgnoreEntries), 1)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "Comment")
	assert.Equal(t, result.IgnoreEntries[0].Original, "# Amazing")
	assert.Equal(t, result.IgnoreEntries[0].Value, "Amazing")
}

func TestIgnoreParseEmpty(t *testing.T) {
	result, _ := ParseIgnore([]string{""})

	assert.Equal(t, len(result.IgnoreEntries), 1)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "Empty")
	assert.Equal(t, result.IgnoreEntries[0].Original, "")
	assert.Equal(t, result.IgnoreEntries[0].Value, "")
}

func TestIgnoreParseWhitespace(t *testing.T) {
	result, _ := ParseIgnore([]string{" "})

	assert.Equal(t, len(result.IgnoreEntries), 1)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "Empty")
	assert.Equal(t, result.IgnoreEntries[0].Original, " ")
	assert.Equal(t, result.IgnoreEntries[0].Value, " ")
}

func TestIgnoreParseWhitespaceBeforePath(t *testing.T) {
	result, _ := ParseIgnore([]string{" .idea"})

	assert.Equal(t, len(result.IgnoreEntries), 1)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "Path")
	assert.Equal(t, result.IgnoreEntries[0].Original, " .idea")
	assert.Equal(t, result.IgnoreEntries[0].Value, " .idea")
}

func TestIgnoreParseMultiLine(t *testing.T) {
	result, _ := ParseIgnore([]string{"# This is a .gitignore file", ".idea", "!foo"})

	assert.Equal(t, len(result.IgnoreEntries), 3)

	assert.Equal(t, result.IgnoreEntries[0].Kind, "Comment")
	assert.Equal(t, result.IgnoreEntries[0].Original, "# This is a .gitignore file")
	assert.Equal(t, result.IgnoreEntries[0].Value, "This is a .gitignore file")

	assert.Equal(t, result.IgnoreEntries[1].Kind, "Path")
	assert.Equal(t, result.IgnoreEntries[1].Original, ".idea")
	assert.Equal(t, result.IgnoreEntries[1].Value, ".idea")

	assert.Equal(t, result.IgnoreEntries[2].Kind, "NegatedPath")
	assert.Equal(t, result.IgnoreEntries[2].Original, "!foo")
	assert.Equal(t, result.IgnoreEntries[2].Value, "foo")
}

func TestRightSpaceNoWhiteSpace(t *testing.T) {
	assert.Equal(t, "foo", TrimRightSpace("foo"))
}

func TestRightSpaceTrailingWhiteSpace(t *testing.T) {
	assert.Equal(t, "bar", TrimRightSpace("bar  "))
}

func TestRightSpaceTrailingEscapedWhiteSpace(t *testing.T) {
	assert.Equal(t, "baz\\ ", TrimRightSpace("baz\\ "))
}

func TestUnescapeNoEscapedChars(t *testing.T) {
	assert.Equal(t, "foo", Unescape("foo"))
}

func TestUnescapeNoEscapedHash(t *testing.T) {
	assert.Equal(t, "#foo", Unescape("\\#foo"))
}

func TestUnescapeEscapedSlashes(t *testing.T) {
	assert.Equal(t, "bar\\", Unescape("bar\\\\"))
}
