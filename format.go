package tools

import "strings"

// Format represents a valid serialization format.
type Format string

const (
	Plain Format = "plain"
	JSON  Format = "json"
	YAML  Format = "yaml"
)

// String returns a string based on the given format.
func (p Format) String() string {
	return string(p)
}

// FormatFromString instances a new Format from the given string.
func FormatFromString(format string) Format {
	return Format(strings.ToLower(format))
}
