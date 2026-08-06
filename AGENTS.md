# AGENTS Guide

This file provides repository-specific guidance for coding agents working in `rakutentech/passenger-go-exporter`.

## Repository Overview

- Language: Go (target version: 1.23.1)
- Purpose: Prometheus exporter for Passenger application metrics
- Entry point: `main.go`
- Core packages:
  - `logging/`
  - `metric/`
  - `passenger/`
- Tests:
  - Unit/integration tests across packages
  - Optional e2e tests in `test/e2e/`

## Local Development Workflow

Use the commands documented in `docs/development.md`:

```bash
go build .
go fmt ./...
golangci-lint run
go test -v ./...
```

For full coverage report:

```bash
go test -coverprofile=cover.out -cover ./... && go tool cover -html=cover.out -o cover.html
```

Optional e2e tests require environment setup (ensure the Passenger instance registry directory is available at `/sock`):

    E2E=true go test ./test/e2e/

## Change Guidelines

- Keep changes minimal and focused on the requested task.
- Add or update tests when behavior changes.
- Do not modify unrelated code.
- Follow existing code and package structure.
- Use `go fmt` formatting standards.

## Documentation and Commit Conventions

- Update docs when behavior, usage, or operations change.
- Prefer Conventional Commit style as documented in `CONTRIBUTING.md`:
  - `docs(...)`, `fix(...)`, `feat(...)`, `refactor(...)`, etc.

## Notes for Agents

- Check `README.md` for runtime usage and deployment examples.
- Check `docs/development.md` for development and testing expectations.
- Avoid running e2e/passenger-dependent tests unless required environment is available.
