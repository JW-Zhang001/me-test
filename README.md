# me-test

## Description
This is the blockchain(CLI) automated testing framework

Note: Requires Go 1.20+

## Installation
```shell
go mod download
go mod tidy
```

## Quick Start
```shell
go run main.go
or 
go test ./... -json | go-test-report -o report/my-test-report.html   
```