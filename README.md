# go-gitignore

A simple gitignore format parser.

This package parses .gitignore files into a publicly exposed struct for usage in other
packages. It does not provide any matching capabilities, nor does it support
marshalling the gitignore Entries.

## API

```
type Entry struct {
  Kind     string
  Value    string
  Original string
}
```

The Parse functions return an array of Entry structs (or an error).

* Kind indicates the type of the entry, either Empty, Comment, Path, or NegatedPath
* Value indicates the effective value of the entry, with the following transformations
  Empty lines are always empty!
  Comments are left trimmed for whitespace
  Paths are unescaped
  NegatedPaths are unescaped with the negation character removed
* Original provides the original line, unaltered

## Example Usage

```go
package main

import (
  "fmt"

  ignore "github.com/shteou/go-gitignore"
)

func main() {
  lines := []string{"# A comment", ".ignore*", "!ignore.dont"}
  ignore, _ := ignore.ParseIgnoreLines(lines)
  fmt.Printf("%v+\n", ignore)
  fmt.Printf("Line 0: Type '%s', Value '%s', Original '%s'\n", ignore.Entries[0].Kind, ignore.Entries[0].Value, ignore.Entries[0].Original)
}
```
