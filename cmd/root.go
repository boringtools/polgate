package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	inputFilePath          string
	inputPolicyFilePath    string
	inputExceptionFilePath string
	workDirectory          string
	policies               []string
	exceptions             []string
)

var rootCmd = &cobra.Command{
	Use:   "PolGate",
	Short: "Security policy and exception management tool",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFilePath, "input", "i", "", "input file path")
	rootCmd.PersistentFlags().StringSliceVarP(&policies, "policy", "p", []string{}, "List of policy")
	rootCmd.PersistentFlags().StringSliceVarP(&exceptions, "exceptions", "e", []string{}, "List of exception")

	rootCmd.PersistentFlags().StringVarP(&workDirectory, "directory", "d", "/tmp/", "working directory path")
	rootCmd.PersistentFlags().StringVarP(&inputPolicyFilePath, "policy-file", "o", "", "policy file path")
	rootCmd.PersistentFlags().StringVarP(&inputExceptionFilePath, "exception-file", "x", "", "exception file path")

	rootCmd.MarkPersistentFlagRequired("input")
	rootCmd.MarkFlagsOneRequired("policy", "policy-file")
}
