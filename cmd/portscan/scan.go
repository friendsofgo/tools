package main

import "github.com/friendsofgo/tools"

func doPortScan(address string, protocol tools.Protocol, from, to int, fmt tools.Format) error {
	formatter, err := tools.NewScanResultsFormatter(fmt)
	if err != nil {
		return err
	}

	results, err := tools.ScanPortRange(protocol, address, from, to)
	if err != nil {
		return err
	}

	var w tools.Writer
	if resultsOutput == emptyString {
		w = tools.NewStdWriter()
	} else {
		w = tools.NewFileWriter(resultsOutput)
	}

	return w.Write(formatter.Format(results))
}
