package main

import (
	"github.com/friendsofgo/tools"
	"github.com/spf13/cobra"
)

const (
	outputFlag = "file"
	formatFlag = "fmt"

	emptyString = ""
)

var resultsOutput string
var resultsFormat string

var rootCmd = &cobra.Command{
	Use:   "portscan",
	Short: "Portscan is a cool port scanner.",
	Long: `
A fast and cool port scanner built with love by Friends of Go.
Complete documentation is available at https://github.com/friendsofgo/tools.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		formatter, err := tools.NewScanResultsFormatter(tools.FormatFromString(resultsFormat))
		if err != nil {
			return err
		}

		results, err := tools.ScanPortRange(tools.TCP, "127.0.0.1", 80, 800)
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
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&resultsOutput, outputFlag, "f", emptyString, "results file (stdin by default)")
	rootCmd.PersistentFlags().StringVar(&resultsFormat, formatFlag, tools.Plain.String(), "results format")
}
