package config

import "github.com/boringtools/polgate/pkg/common"

func GetFilepaths() map[string]string {
	FilePaths := map[string]string{
		"semgrep_policy":    common.WorkDirectory + "/semgrep_policy.json",
		"semgrep_exception": common.WorkDirectory + "/semgrep_exception.json",
	}

	return FilePaths
}
