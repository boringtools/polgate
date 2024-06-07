package policy

import (
	"encoding/json"

	"github.com/boringtools/polgate/pkg/common"
	"github.com/boringtools/polgate/pkg/config"
	"github.com/boringtools/polgate/pkg/logger"
	"github.com/boringtools/polgate/pkg/models"
)

var (
	sgPFailSeverity models.SGPFailSeverity
)

func GenSemGrepPolicy() {
	errUnMarshal := json.Unmarshal([]byte(common.SGPFailSeverityTemp), &sgPFailSeverity)

	if errUnMarshal != nil {
		logger.LogERRP("error[GenSemGrepPolicy] : ", errUnMarshal)
	}

	if len(common.PolicyList) != 0 {

		for _, value := range common.PolicyList {
			policyData := PolicyMapping()[value]
			sgPFailSeverity.Semgrep.FailOn = append(sgPFailSeverity.Semgrep.FailOn, policyData...)
		}
	}

	if common.InputPolicyFilePath != "" {

		for _, value := range policyFile.Policies {
			policyData := PolicyMapping()[value]
			sgPFailSeverity.Semgrep.FailOn = append(sgPFailSeverity.Semgrep.FailOn, policyData...)
		}
	}

	jsonMarshal, errMarshal := json.Marshal(sgPFailSeverity)

	if errMarshal != nil {
		logger.LogERRP("error[GenSemGrepPolicy] : ", errMarshal)
	}

	errSaveToJSON := common.SaveToJSONFile(config.GetFilepaths()["semgrep_policy"], jsonMarshal)

	if errSaveToJSON != nil {
		logger.LogERRP("error[GenSemGrepPolicy] : ", errSaveToJSON)
	}
}
