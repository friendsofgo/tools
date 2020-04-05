package main

import (
	"errors"
	"github.com/friendsofgo/tools"
	"github.com/spf13/cobra"
	"strconv"
)

const (
	outputFlag = "file"
	formatFlag = "fmt"

	emptyString = ""

	singleNumArgs = 3
	rangeNumArgs  = 4
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
}

var cmdSingle = &cobra.Command{
	Use:   "single ADDRESS PROTOCOL PORT",
	Short: "Does a single port scan (i.e. 8080)",
	Long: `
Does a single scan for the given port on the given address through the given protocol.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(singleNumArgs)(cmd, args); err != nil {
			return err
		}

		if !tools.IsAllowedProtocol(tools.ProtocolFromString(args[1])) {
			return errors.New("invalid transport layer protocol - only tcp or udp allowed")
		}

		if !tools.IsNumber(args[2]) {
			return errors.New("invalid port format - must be a number")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		address := args[0]
		protocol := tools.ProtocolFromString(args[1])
		port, _ := strconv.Atoi(args[2])
		format := tools.FormatFromString(resultsFormat)

		return doPortScan(address, protocol, port, port, format)
	},
}

var cmdRange = &cobra.Command{
	Use:   "range ADDRESS PROTOCOL FROM TO",
	Short: "Does a range port scan (i.e. 0-1024)",
	Long: `
Does a range scan for each port number of the given range on the given address through the given protocol.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(rangeNumArgs)(cmd, args); err != nil {
			return err
		}

		if !tools.IsAllowedProtocol(tools.ProtocolFromString(args[1])) {
			return errors.New("invalid transport layer protocol - only tcp or udp allowed")
		}

		if !tools.IsNumber(args[2]) || !tools.IsNumber(args[3]) {
			return errors.New("invalid port format - must be a number")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		address := args[0]
		protocol := tools.ProtocolFromString(args[1])
		from, _ := strconv.Atoi(args[2])
		to, _ := strconv.Atoi(args[3])
		format := tools.FormatFromString(resultsFormat)

		return doPortScan(address, protocol, from, to, format)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&resultsOutput, outputFlag, "f", emptyString, "results file (stdin by default)")
	rootCmd.PersistentFlags().StringVar(&resultsFormat, formatFlag, tools.Plain.String(), "results format, allowed: plain, json, yaml")

	rootCmd.AddCommand(cmdSingle)
	rootCmd.AddCommand(cmdRange)
}
