package common

var (
	SGPFailSeverityTemp  string
	SemGrepExceptionTemp string
)

func init() {
	SGPFailSeverityTemp = `
	{
		"semgrep": {
			"fail_on": []
		}
	}`

	SemGrepExceptionTemp = `
	{
		"exceptions": {
			"fingerprint": []
		}
	}`
}
