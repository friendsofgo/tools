package tools

// Protocol represents a valid transport layer protocol.
type Protocol string

const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)

func (p Protocol) String() string {
	return string(p)
}
