# gototal-cobertura

A Go utility that reports the total line-rate from a cobertura XML code coverage report.

## Usage

This is useful in cases like the [gitlab pipelines](https://docs.gitlab.com/ee/ci/pipelines/settings.html#add-test-coverage-results-to-a-merge-request) that take a regular expression and uses the last match found from the output to find the coverage.  For `go test --cover` this becomes an issue if you have multiple packages as only the last packages coverage will be reported on.

This utility will read the generated cobertura report and output the total coverage.  If this utility is run at the end of the gitlab job, then the regex can match on the total code coverage.

An example `.gitlab-ci.yml` job might look something like the following:

``` yml
golang:
  stage: test
  script:
    - go install gotest.tools/gotestsum
    - go install github.com/boumenot/gocover-cobertura
    - go install github.com/ggere/gototal-cobertura
    - gotestsum --junitfile report.xml --format testname -- -coverprofile=coverage.out ./...
    - gocover-cobertura < coverage.out > coverage.xml
    - gototal-cobertura < coverage.xml
  artifacts:
    when: always
    reports:
      junit: report.xml
      cobertura: coverage.xml
```

## Output

```sh
Total coverage: 62.50%
```

## Installation

```sh
go install gitlab.com/ggere/gototal-cobertura@latest
```
