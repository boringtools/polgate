package policy

import (
	"encoding/json"
	"os"

	"github.com/boringtools/polgate/pkg/common"
	"github.com/boringtools/polgate/pkg/logger"
	"github.com/boringtools/polgate/pkg/models"
)

var (
	policyFile models.PolicyFile
	inputFile  interface{}
)

func IsPolicyExists(policyName string) {
	allPolicies := PolicyMapping()

	_, exists := allPolicies[policyName]

	if !exists {
		logger.LogERRP("Invalid policy, please check list of supported policy : ", policyName)
		os.Exit(1)
	}
}

func ValidatePolicyFile() error {
	jsonContent, errContent := common.GetJSONFileContent(common.InputPolicyFilePath)

	if errContent != nil {
		return errContent
	}

	errUnMarshal := json.Unmarshal(jsonContent, &policyFile)

	if errUnMarshal != nil {
		return errUnMarshal
	}

	return nil
}

func ValidatePolicy() {

	for _, value := range common.PolicyList {
		IsPolicyExists(value)
	}

	if common.InputPolicyFilePath != "" {

		errValidatPolicyFile := ValidatePolicyFile()

		if errValidatPolicyFile != nil {
			logger.LogERRP("error[ValidatePolicy] : ", errValidatPolicyFile)
		}

		for _, value := range policyFile.Policies {
			IsPolicyExists(value)
		}
	}
}

func ValidateInputFile() {

	checkInputFile, errFile := common.GetJSONFileContent(common.InputFilePath)

	if errFile != nil {
		logger.LogERRP("error[ValidateInputFile] : ", errFile)
	}

	errUnMarshall := json.Unmarshal(checkInputFile, &inputFile)

	if errUnMarshall != nil {
		logger.LogERRP("error[ValidateInputFile] : ", errUnMarshall)
	}
}
