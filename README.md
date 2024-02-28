<p align="center">
    <a href="https://github.com/OpsLevel/opslevel-jq-parser/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/OpsLevel/opslevel-jq-parser.svg" alt="License" /></a>
    <a href="https://go.dev">
        <img src="https://img.shields.io/github/go-mod/go-version/OpsLevel/opslevel-jq-parser" alt="Made With Go" /></a>
    <a href="https://GitHub.com/OpsLevel/opslevel-jq-parser/releases/">
        <img src="https://img.shields.io/github/v/release/OpsLevel/opslevel-jq-parser?include_prereleases" alt="Release" /></a>
    <a href="https://GitHub.com/OpsLevel/opslevel-jq-parser/issues/">
        <img src="https://img.shields.io/github/issues/OpsLevel/opslevel-jq-parser.svg" alt="Issues" /></a>
    <a href="https://github.com/OpsLevel/opslevel-jq-parser/graphs/contributors">
        <img src="https://img.shields.io/github/contributors/OpsLevel/opslevel-jq-parser" alt="Contributors" /></a>
    <a href="https://github.com/OpsLevel/opslevel-jq-parser/pulse">
        <img src="https://img.shields.io/github/commit-activity/m/OpsLevel/opslevel-jq-parser" alt="Activity" /></a>
    <a href="https://codecov.io/gh/OpsLevel/opslevel-jq-parser">
        <img src="https://codecov.io/gh/OpsLevel/opslevel-jq-parser/branch/main/graph/badge.svg" alt="CodeCov" /></a>
    <a href="https://dependabot.com/">
        <img src="https://badgen.net/badge/Dependabot/enabled/green?icon=dependabot" alt="Dependabot" /></a>
    <a href="https://pkg.go.dev/github.com/opslevel/opslevel-jq-parser/v2024">
        <img src="https://pkg.go.dev/badge/github.com/opslevel/opslevel.svg" alt="Go Reference" /></a>
</p>

[![Overall](https://img.shields.io/endpoint?style=flat&url=https%3A%2F%2Fapp.opslevel.com%2Fapi%2Fservice_level%2FAN4c4UlHKKLbrHAlFzF4FKXpeGYnjEtC5765UYF1Exc)](https://app.opslevel.com/services/opslevel-jq-parser/maturity-report)



# opslevel-jq-parser
A jq wrapper which aids in converting data to opslevel-go input structures

This library leverages https://github.com/flant/libjq-go which are CGO bindings to the JQ library which provide C native speed

#  Installation

```bash
go get github.com/opslevel/opslevel-jq-parser/v2024
```

Then wherever you compile or test that project you'll need to add

```bash
docker run --name "libjq" -d flant/jq:b6be13d5-glibc
docker cp libjq:/libjq ./libjq 
docker rm libjq
export CGO_ENABLED=1
export CGO_CFLAGS="-I$(pwd)/libjq/include"
export CGO_LDFLAGS="-L$(pwd)/libjq/lib"
```

Here is a nice stanza you can put into your GitHub Actions workflow files

> NOTE: the version is important - please see https://github.com/flant/libjq-go#notes

```yaml
      - name: Setup LibJQ
        run: |-
          docker run --name "libjq" -d flant/jq:b6be13d5-glibc
          docker cp libjq:/libjq ./libjq 
          docker rm libjq
          echo CGO_ENABLED=1 >> $GITHUB_ENV
          echo CGO_CFLAGS="-I$(pwd)/libjq/include" >> $GITHUB_ENV
          echo CGO_LDFLAGS="-L$(pwd)/libjq/lib" >> $GITHUB_ENV
```