package policy

import (
	"context"

	"github.com/boringtools/polgate/pkg/common"
	"github.com/boringtools/polgate/pkg/logger"
	"github.com/open-policy-agent/opa/rego"
)

func EvalPolicy() (status bool, reason []string) {

	ctx := context.Background()

	r := rego.New(
		rego.Query("x = data.polgate"),
		rego.Load([]string{common.WorkDirectory}, nil),
		rego.Module("default.rego", string(common.RegoDefault)),
		rego.Module("semgrep_fail_severity.rego", string(common.RegoSGFailSeverity)),
	)

	query, errPrep := r.PrepareForEval(ctx)

	if errPrep != nil {
		logger.LogERRP("error[EvalPolicy] : ", errPrep)
	}

	evalResult, errEval := query.Eval(ctx, rego.EvalInput(inputFile))

	if errEval != nil {
		logger.LogERRP("error[EvalPolicy] : ", errEval)
	}

	data := evalResult[0].Bindings["x"]
	convertData := data.(map[string]interface{})

	checkAllow := convertData["allow"]
	checkViolations := convertData["violations"].(map[string]interface{})

	for key := range checkViolations {
		reason = append(reason, key)
	}

	return checkAllow.(bool), reason
}
