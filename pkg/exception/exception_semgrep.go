package exception

import (
	"encoding/json"

	"github.com/boringtools/polgate/pkg/common"
	"github.com/boringtools/polgate/pkg/config"
	"github.com/boringtools/polgate/pkg/logger"
	"github.com/boringtools/polgate/pkg/models"
)

var (
	semgrepException models.SemgrepException
)

func GenSemGrepException() {
	errUnMarshal := json.Unmarshal([]byte(common.SemGrepExceptionTemp), &semgrepException)

	if errUnMarshal != nil {
		logger.LogERRP("error[CheckSemGrepExceptionTemp] : ", errUnMarshal)
	}

	if len(common.ExceptionList) != 0 {
		semgrepException.Exceptions.Fingerprint = append(semgrepException.Exceptions.Fingerprint, common.ExceptionList...)
	}

	if common.InputExceptionFilePath != "" {
		semgrepException.Exceptions.Fingerprint = append(semgrepException.Exceptions.Fingerprint, exceptionFile.Exceptions...)
	}

	jsonMarshal, errMarshal := json.Marshal(semgrepException)

	if errMarshal != nil {
		logger.LogERRP("error[GenSemGrepException] : ", errMarshal)
	}

	errSaveToJSON := common.SaveToJSONFile(config.GetFilepaths()["semgrep_exception"], jsonMarshal)

	if errSaveToJSON != nil {
		logger.LogERRP("error[GenSemGrepException] : ", errSaveToJSON)
	}
}
