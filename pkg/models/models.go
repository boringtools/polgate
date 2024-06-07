package models

type SGPFailSeverity struct {
	Semgrep struct {
		FailOn []string `json:"fail_on"`
	} `json:"semgrep"`
}

type SemgrepException struct {
	Exceptions struct {
		Fingerprint []string `json:"fingerprint"`
	} `json:"exceptions"`
}

type PolicyFile struct {
	Policies []string `json:"policies"`
}

type ExceptionFile struct {
	Exceptions []string `json:"exceptions"`
}
