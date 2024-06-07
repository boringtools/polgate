package polgate

import rego.v1

violations[msg] if {
    some issue in input.results
    issue.extra.severity = data.semgrep.fail_on[_]
    not semgrep_exceptions(issue)

    msg := sprintf("SemGrep result have issues with %s", [data.semgrep.fail_on[_]])
}

semgrep_exceptions(issue) if {
    data.exceptions.fingerprint[_] = issue.extra.fingerprint
}
