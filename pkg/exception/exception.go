package exception

import (
	"encoding/json"

	"github.com/boringtools/polgate/pkg/common"
	"github.com/boringtools/polgate/pkg/models"
)

var (
	exceptionFile models.ExceptionFile
)

func ValidateExceptionFile() error {
	jsonContent, errContent := common.GetJSONFileContent(common.InputExceptionFilePath)

	if errContent != nil {
		return errContent
	}

	errUnMarshal := json.Unmarshal(jsonContent, &exceptionFile)

	if errUnMarshal != nil {
		return errUnMarshal
	}

	return nil
}
