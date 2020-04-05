package tools

import "strings"

// Protocol represents a valid transport layer protocol.
type Protocol string

const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)

// String returns a string based on the given protocol.
func (p Protocol) String() string {
	return string(p)
}

// ProtocolFromString instances a new Protocol from the given string.
func ProtocolFromString(protocol string) Protocol {
	return Protocol(strings.ToLower(protocol))
}