package reporter

import (
	"os"

	"github.com/boringtools/polgate/pkg/logger"
	"github.com/boringtools/polgate/pkg/policy"
	"github.com/jedib0t/go-pretty/v6/table"
)

func GenOutput() {

	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"Reason"})

	status, reasons := policy.EvalPolicy()

	if status {
		logger.Log("Policy Gate : Pass [No violations found]")
		os.Exit(0)
	} else {
		logger.LogERR("Policy Gate : Fail [Policy violations found]")

		for _, value := range reasons {
			tbl.AppendRows([]table.Row{
				{value},
			})
		}
		tbl.AppendSeparator()
		tbl.Render()
		os.Exit(1)
	}
}
