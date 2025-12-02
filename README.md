# Go Project

Small, idiomatic Go repository template with commands for building, testing, linting, and releasing.

## Prerequisites

- Go 1.20+ installed
- git

## Getting started

1. Clone
    git clone <repo-url>
2. Enter
    cd <repo-dir>
3. Initialize module (if not present)
    go mod init github.com/username/project-name

## Build

- Build all packages:
  go build ./...
- Build a command in `cmd/<name>`:
  go build -o bin/<name> ./cmd/<name>

## Run

- Run a single package:
  go run ./cmd/<name>
- Run tests:
  go test ./... -v

## Formatting, vetting, and linting

- Format:
  go fmt ./...
- Vet:
  go vet ./...
- Static analysis (recommended):
  golangci-lint run


## Contributing

- Open issues or PRs.
- Follow gofmt and run tests locally before submitting.
- Add tests for any new behavior.

## License

Add a LICENSE file (e.g., MIT, Apache 2.0) and reference it here.

---

Replace placeholders (repo URL, command names) with project-specific values.