package markup

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	GenerationWarning = "Tcl Module Generated by tcltm; DO NOT EDIT"
)

// Write the string to the buffer
func Write(b *bytes.Buffer, s string) (int, error) {
	return b.WriteString(s)
}

// WriteLine will write the string to the buffer with a newline <LF>.
func WriteLine(b *bytes.Buffer, s string) (int, error) {
	return Write(b, fmt.Sprintf("%s\n", s))
}

// Divider will write a divider line to the buffer
func Divider(b *bytes.Buffer) (int, error) {
	return Write(b, fmt.Sprintf("# %s\n", strings.Repeat("#", 78)))
}

// Comment will write the string to the buffer as a comment
func Comment(b *bytes.Buffer, s string) (int, error) {
	return Write(b, fmt.Sprintf("# %s", s))
}

// Commentln will write the string to the buffer as a comment
// with a newline <LF>
func Commentln(b *bytes.Buffer, s string) (int, error) {
	return Comment(b, fmt.Sprintf("%s\n", s))
}

// NewLine will write a newline to the buffer
func NewLine(b *bytes.Buffer) (int, error) {
	return Write(b, "\n")
}

// Meta will write a meta tag to the buffer
// key is forced to upper case
func Meta(b *bytes.Buffer, key, value string) (int, error) {
	return Comment(b, fmt.Sprintf("%s: %s\n", strings.ToUpper(key), value))
}

// EOF