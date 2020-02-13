module todo/tools

go 1.11

require (
	github.com/golang/mock v1.4.0
	github.com/golangci/golangci-lint v1.23.6
	github.com/kamilsk/egg v0.0.13
	golang.org/x/tools v0.3.2
)

replace github.com/izumin5210/gex => github.com/kamilsk/gex v0.6.0-e4

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.3
