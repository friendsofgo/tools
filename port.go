package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"net"
	"sync"
	"time"
)

const (
	MinValidPortNumber = 0
	MaxValidPortNumber = 65535

	PortScanTimeout = 60 * time.Second
)

func isValidPort(port int) bool {
	return MinValidPortNumber <= port && port <= MaxValidPortNumber
}

// ScanPort does a Dial check of the given port at the given address
// through the given protocol. It returns an ScanResult with the info.
func ScanPort(protocol Protocol, hostname string, port int) (ScanResult, error) {
	if !isValidPort(port) {
		return ScanResult{}, errors.New("invalid port - out of range")
	}

	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol.String(), address, PortScanTimeout)
	if err != nil {
		return ScanResult{Port: port, State: Closed, Protocol: protocol}, nil
	}

	_ = conn.Close()
	return ScanResult{Port: port, State: Open, Protocol: protocol}, nil
}

// ScanPortRange does a concurrent ScanPort call for each port present
// within the given ports range. It returns a slide of ScanResults with the info.
func ScanPortRange(protocol Protocol, hostname string, from, to int) (results []ScanResult, err error) {
	if !isValidPort(from) || !isValidPort(to) {
		return nil, errors.New("invalid port - out of range")
	}

	ch := make(chan ScanResult)
	wg := &sync.WaitGroup{}

	go func(ch <-chan ScanResult) {
		for result := range ch {
			results = append(results, result)
			wg.Done()
		}
	}(ch)

	for i := from; i <= to; i++ {
		port := i
		wg.Add(1)

		go func(ch chan<- ScanResult) {
			result, _ := ScanPort(protocol, hostname, port)
			ch <- result
		}(ch)
	}

	wg.Wait()
	close(ch)

	return
}

// ScanResult contains all the information related with an specific port scan.
// It includes the protocol, the port and the state {open, closed}.
type ScanResult struct {
	Protocol Protocol `json:"protocol"`
	Port     int      `json:"port"`
	State    State    `json:"state"`
}

// State represents the state of a given port {open, closed}.
type State string

const (
	Open   State = "open"
	Closed       = "closed"
)

// ScanResultFormatter defines the behaviour needed to format
// a given slide of ScanResults into the formatter format.
type ScanResultsFormatter interface {
	Format(results []ScanResult) []byte
}

type jsonScanResultsFormatter struct{}

// NewJSONScanResultsFormatter returns a new JSON results formatter.
func NewJSONScanResultsFormatter() *jsonScanResultsFormatter {
	return &jsonScanResultsFormatter{}
}

func (*jsonScanResultsFormatter) Format(results []ScanResult) (bytes []byte) {
	bytes, _ = json.Marshal(results)
	return bytes
}

type yamlScanResultsFormatter struct{}

// NewYAMLScanResultsFormatter returns a new YAML results formatter.
func NewYAMLScanResultsFormatter() *yamlScanResultsFormatter {
	return &yamlScanResultsFormatter{}
}

func (*yamlScanResultsFormatter) Format(results []ScanResult) (bytes []byte) {
	bytes, _ = yaml.Marshal(results)
	return bytes
}

type plainScanResultsFormatter struct{}

// NewPlainScanResultsFormatter returns a new plain results formatter.
func NewPlainScanResultsFormatter() *plainScanResultsFormatter {
	return &plainScanResultsFormatter{}
}

func (*plainScanResultsFormatter) Format(results []ScanResult) (bytes []byte) {
	for _, result := range results {
		out := fmt.Sprintf("Protocol: %s    Port: %d    State: %s\n", result.Protocol, result.Port, result.State)
		bytes = append(bytes, []byte(out)...)
	}

	return
}

// NewScanResultsFormatter is a factory method that returns a
// valid ScanResultsFormatter implementation or an error.
func NewScanResultsFormatter(format Format) (ScanResultsFormatter, error) {
	switch format {
	case Plain:
		return NewPlainScanResultsFormatter(), nil
	case JSON:
		return NewJSONScanResultsFormatter(), nil
	case YAML:
		return NewYAMLScanResultsFormatter(), nil
	default:
		return nil, errors.New("non-valid results output")
	}
}
