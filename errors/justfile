bash := if os() == "windows" { "bash.exe" } else { "/usr/bin/env bash" }

# Builds the library
build:
    go build -o ./bin/errors

# Runs the tests
test:
    go test -v ./...

# Runs the tests and coverage
cov:
    #!{{ bash }}
    go test -coverprofile=./artifacts/coverage.out -covermode=atomic ./...
    go tool cover -html=./artifacts/coverage.out -o ./artifacts/coverage/coverage.html

# Runs the tests and coverage and opens the report
cov-open:
    go test -coverprofile=./artifacts/coverage.out -covermode=atomic ./... && go tool cover -html=coverage.out

# Formats the code
fmt:
    go fmt ./...

# Lints the code
lint:
    golangci-lint run ./...

# Packs the library, runs go mod tidy
pack:
    go mod tidy