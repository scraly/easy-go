# GolangCI Lint

https://golangci-lint.run/usage/linters/

## shadow

```go
linters-settings:
  govet:
    # Report about shadowed variables.
    # Default: false
    check-shadowing: true
    # Settings per analyzer.
    settings:
      shadow:
        # Whether to be strict about shadowing; can be noisy.
        # Default: false
        strict: true
```