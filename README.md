# PolGate

`PolGate` is a tool for security policies and exception management. Polgate helps you to enable effective gatekeeping when integrated in the CI/CD pipeline.

## Getting started

> Ensure $(go env GOPATH)/bin is in your $PATH

```bash
go install github.com/boringtools/polgate@main
```

## Supported tools

- [SemGrep](https://github.com/semgrep/semgrep)

> For more tools, Please create a issue.

## Usage

Fail workflow when SemGrep results in `error`(high) severity findings.

```bash
polgate eval --input /input/file/path/results.json --policy semgrep_fail_error
```

Add exceptions or false positives to continue the workflow.

```bash
polgate eval --input /input/file/path/results.json --policy semgrep_fail_error --exceptions exceptions,seprated,by,comma
```

Apply list of supported policies form a JSON file

```bash
polgate eval --input /input/file/path/results.json --policy-file /policy/file/path/policies.json
```

Apply list of exceptions form a JSON file

```bash
polgate eval --input /input/file/path/results.json --policy semgrep_fail_error --exception-file /exceptions/file/path/exceptions.json
```

## List of supported policies

- semgrep_pass_all
- semgrep_fail_error
- semgrep_fail_error_warning
- semgrep_fail_all

> For more policies, Please create a issue.

## Sample workflows

- [Workflow fail with policy gate](https://github.com/c0d3G33k/vulnado01/actions/runs/9469776634/job/26089218585)
- [Workflow pass with policy gate](https://github.com/c0d3G33k/dvna/actions/runs/9469857662/job/26089475901)
  
## Policy JSON file template

```json
{
    "policies": [
        "semgrep_fail_all"
    ]
}
```

## Exception JSON file template

```json
{
    "exceptions": [
        "sample-fingerprint",
    ]
}
```
