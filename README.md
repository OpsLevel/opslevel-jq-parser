<p align="center">
    <a href="https://github.com/OpsLevel/opslevel-jq-parser/blob/main/LICENSE" alt="License">
        <img src="https://img.shields.io/github/license/OpsLevel/opslevel-jq-parser.svg" /></a>
    <a href="http://golang.org" alt="Made With Go">
        <img src="https://img.shields.io/github/go-mod/go-version/OpsLevel/opslevel-jq-parser" /></a>
    <a href="https://GitHub.com/OpsLevel/opslevel-jq-parser/releases/" alt="Release">
        <img src="https://img.shields.io/github/v/release/OpsLevel/opslevel-jq-parser?include_prereleases" /></a>
    <a href="https://GitHub.com/OpsLevel/opslevel-jq-parser/issues/" alt="Issues">
        <img src="https://img.shields.io/github/issues/OpsLevel/opslevel-jq-parser.svg" /></a>
    <a href="https://github.com/OpsLevel/opslevel-jq-parser/graphs/contributors" alt="Contributors">
        <img src="https://img.shields.io/github/contributors/OpsLevel/opslevel-jq-parser" /></a>
    <a href="https://github.com/OpsLevel/opslevel-jq-parser/pulse" alt="Activity">
        <img src="https://img.shields.io/github/commit-activity/m/OpsLevel/opslevel-jq-parser" /></a>
	<a href="https://codecov.io/gh/OpsLevel/opslevel-jq-parser">
  		<img src="https://codecov.io/gh/OpsLevel/opslevel-jq-parser/branch/main/graph/badge.svg"/></a>
    <a href="https://dependabot.com/" alt="Dependabot">
        <img src="https://badgen.net/badge/Dependabot/enabled/green?icon=dependabot" /></a>
    <a href="https://pkg.go.dev/github.com/opslevel/opslevel-jq-parser/v2024" alt="Go Reference">
        <img src="https://pkg.go.dev/badge/github.com/opslevel/opslevel.svg" /></a>
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

Here is a nice stanza you can put into your github actions workflow files

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