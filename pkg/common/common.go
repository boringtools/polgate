package common

import (
	"embed"
	"io"
	"os"

	"github.com/boringtools/polgate/pkg/logger"
)

var (
	InputFilePath          string
	InputPolicyFilePath    string
	InputExceptionFilePath string
	WorkDirectory          string
	PolicyList             []string
	ExceptionList          []string
	RegoDefault            []byte
	RegoSGFailSeverity     []byte
	RegoPolicyDir          embed.FS
)

func GetJSONFileContent(filePath string) (byteData []byte, err error) {
	locateFile, errLocateFile := os.Open(filePath)

	if errLocateFile != nil {
		return nil, errLocateFile
	}

	byteData, errRead := io.ReadAll(locateFile)

	if errRead != nil {
		return nil, errRead
	}

	return byteData, nil
}

func SaveToJSONFile(filePath string, content []byte) error {
	errFileWrite := os.WriteFile(filePath, content, 0644)

	if errFileWrite != nil {
		return errFileWrite
	}

	return nil
}

func SetupDirectory() {
	WorkDirectory = WorkDirectory + "polgate"

	_, errDirExists := os.Stat(WorkDirectory)

	if errDirExists != nil {
		errDir := os.Mkdir(WorkDirectory, 0755)

		if errDir != nil {
			logger.LogERRP("error[SetupDirectory] : ", errDir)
		}
	}
}

func RegoPolicyFiles() {
	defaultRego, errDefault := RegoPolicyDir.ReadFile("internal/policy/default.rego")

	if errDefault != nil {
		logger.LogERRP("error[RegoPolicyFiles]", errDefault)
	}
	semgrep_fail_severity, errSGFailSeverity := RegoPolicyDir.ReadFile("internal/policy/semgrep_fail_severity.rego")

	if errSGFailSeverity != nil {
		logger.LogERRP("error[RegoPolicyFiles]", errDefault)
	}

	RegoDefault = defaultRego
	RegoSGFailSeverity = semgrep_fail_severity
}
