package policy

func PolicyMapping() map[string][]string {

	PolicyNames := make(map[string][]string)

	PolicyNames["semgrep_pass_all"] = []string{""}
	PolicyNames["semgrep_fail_error"] = []string{"ERROR"}
	PolicyNames["semgrep_fail_error_warning"] = []string{"ERROR", "WARNING"}
	PolicyNames["semgrep_fail_all"] = []string{"ERROR", "WARNING", "INFO"}

	return PolicyNames
}

func PolicyMappingRego() map[string]string {

	PolicyNames := make(map[string]string)

	PolicyNames["semgrep_pass_all"] = "semgrep_fail_severity.rego"
	PolicyNames["semgrep_fail_error"] = "semgrep_fail_severity.rego"
	PolicyNames["semgrep_fail_error_warning"] = "semgrep_fail_severity.rego"
	PolicyNames["semgrep_fail_all"] = "semgrep_fail_severity.rego"

	return PolicyNames
}
