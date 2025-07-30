```markdown
# 📚 understanding-algorithms

Repository dedicated to the study and implementation of fundamental algorithms and data structures using the Go programming language. This project serves as a practical exercise book to enhance knowledge in Go and Computer Science, focusing on simplicity and testability.

## ✨ How to Use and Test

### Prerequisites

Ensure you have [Go](https://go.dev/doc/install) installed on your machine (version 1.16 or higher recommended) and `make` (generally pre-installed on Unix/Linux/WSL based systems).

### Clone the Repository

```bash
git clone https://github.com/dmarins/understanding-algorithms.git
cd understanding-algorithms
```

### Configure Go Modules

Ensure all Go dependencies (if any) are resolved:

```bash
go mod tidy
```

### Run Tests

It is highly recommended to run tests to verify the correctness of the implementations. We use a `Makefile` to simplify test execution.

*   **To run ALL tests**:
    ```bash
    make tests
    ```
    (This command executes `go test -cover -p=1 ./...`)

*   **To run ONLY tests from a specific file** (e.g., `binary_search_test.go`):
    ```bash
    make test-binary-search
    ```
    (This command executes `go test -cover -p=1 ./binary-search/`)

## 📄 License (Optional)

This project is licensed under the [MIT License](LICENSE).
```