module tool/tools

go 1.16

require (
	github.com/golang/mock v1.5.0
	github.com/golangci/golangci-lint v1.40.1
	github.com/kyoh86/looppointer v0.1.7
	golang.org/x/exp v0.0.0-20210514180818-737f94c0881e
	golang.org/x/tools v0.1.2-0.20210512205948-8287d5da45e4
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.1.1
