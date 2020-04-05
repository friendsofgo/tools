package tools

// Format represents a valid serialization format.
type Format string

const (
	Plain Format = "plain"
	JSON  Format = "json"
	YAML  Format = "yaml"
)

func (p Format) String() string {
	return string(p)
}

func FormatFromString(format string) Format {
	return Format(format)
}
