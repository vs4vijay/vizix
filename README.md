# vizix

## Getting Started

```console
go get github.com/vs4vijay/vizix
```

---

## Development

---

## Testing

includes following suite of tests.
- `make test-lint`: runs linter/style checks
- `make test-unit`: runs basic unit tests
- `make test`: runs all of the above

---

### Development Notes
```

~/go/bin/cobra init --pkg-name github.com/vs4vijay/vizix

go mod init github.com/vs4vijay/vizix

go build

~/go/bin/cobra add list


gofmt
golint
go vet

git config core.hooksPath .githooks

```