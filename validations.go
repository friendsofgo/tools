package tools

import "strconv"

// IsNumber returns true if the given input can be converted
// into an integer, otherwise returns false.
func IsNumber(n string) bool {
	_, err := strconv.Atoi(n)
	return err == nil
}

// IsValidPort returns true if the given input is
// within the range of valid port numbers, otherwise returns false.
func IsValidPort(port int) bool {
	return MinValidPortNumber <= port && port <= MaxValidPortNumber
}

// IsAllowedProtocol returns true if the given input is a valid
// value for the transport layer protocol. TODO: Use const values.
func IsAllowedProtocol(protocol Protocol) bool {
	return protocol == TCP || protocol == UDP
}
