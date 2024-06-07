package cmd

import (
	"github.com/boringtools/polgate/pkg/common"
	"github.com/boringtools/polgate/pkg/exception"
	"github.com/boringtools/polgate/pkg/policy"
	"github.com/boringtools/polgate/pkg/reporter"
	"github.com/spf13/cobra"
)

var evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Evaluate security policy",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		common.InputFilePath = inputFilePath
		common.WorkDirectory = workDirectory

		common.PolicyList = policies
		common.InputPolicyFilePath = inputPolicyFilePath

		common.ExceptionList = exceptions
		common.InputExceptionFilePath = inputExceptionFilePath

		common.SetupDirectory()
		common.RegoPolicyFiles()

		policy.ValidateInputFile()
		policy.ValidatePolicy()
		exception.ValidateExceptionFile()

		policy.GenSemGrepPolicy()
		exception.GenSemGrepException()

		reporter.GenOutput()
	},
}

func init() {
	rootCmd.AddCommand(evalCmd)
}
