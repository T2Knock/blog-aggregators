# AGENTS

## Build & Test

- Build: `go build -o bin/app .`
- Lint Go: `go fmt ./... && go vet ./...`
- SQL Lint: `sqlfluff lint sql/`
- Run all tests: `go test ./...`
- Run single test: `go test ./handler_user.go -run ^TestName$`

## Code Style Guidelines

- Imports: group stdlib, blank, 3rd-party, blank, internal.
- Formatting: enforce `go fmt` & `goimports`.
- Naming: Exported PascalCase, unexported camelCase, acronyms uppercase.
- Types: interfaces end with "er", avoid unused types.
- Error Handling: check `err` immediately, wrap with `fmt.Errorf`, return upstream.
- Handlers: use `context.Context`, return HTTP codes via standard errors.
- DB: use sqlc-generated code; migrations under `sql/schema` with timestamp prefix.
- Config: use `internal/config`, env vars uppercase with underscores.
- Cursor & Copilot rules: none detected.
- Agents should follow these guidelines.
