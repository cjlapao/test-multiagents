### Task

A standalone Go script that prints "Hello, Agent" to stdout.

### Assigned Specialist

golang-engineer

### Parent Feature

simple-go-hello

### Depends on

none

### Acceptance Criteria

- [x] The Go script compiles without errors (`go build` succeeds).
- [x] Running the executable prints exactly `Hello, Agent` followed by a newline to stdout.
- [x] When run with `go run .`, the program exits with status code 0.

### Definition of Done

- [x] Code implemented following best practices.
- [x] Unit tests written and passing.
- [ ] Reviewed and approved.

### Status

ready-for-review

### Implementation Notes

**Files created:**

| File | Purpose |
|---|---|
| `go.mod` | Go module declaration (`simple-go-hello`) |
| `main.go` | Entry-point package with `run()` logic extracted for testability |
| `main_test.go` | Table-less unit test verifying exact stdout output |

**Key design decisions:**

1. **Extracted `run()` from `main()`** — separates business logic from bootstrapping, enabling clean testing without shelling out or mocking `os.Stdout`.
2. **Exported `outputWriter` variable** — holds an `io.Writer` (defaults to `os.Stdout`). Tests swap it for a `bytes.Buffer`. This avoids cgo-dependent `-race` mocks and keeps the production path untouched.
3. **`fmt.Fprintln(outputWriter, ...)` instead of `fmt.Println(...)`** — ensures whatever writer is assigned receives the output; `fmt.Println` always hits `os.Stdout` regardless.

**Tests added:**

- `TestRun` — replaces `outputWriter` with a `bytes.Buffer`, calls `run()`, asserts the buffer contains `"Hello, Agent\n"`. Restores original writer in a deferred cleanup.

**Verification commands run:**

```bash
go build ./...        # success, no errors
go vet ./...          # success, no issues
gofmt -l .            # empty (all files formatted)
go test ./... -v      # TestRun PASS
go run .              # prints "Hello, Agent", exit code 0
```

CGO-enabled `-race` flag was skipped because no C compiler (`cc`) is available in this environment. No race-prone code exists anyway (single-threaded CLI program).

