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

Commands:

~/go/bin/cobra init --pkg-name github.com/vs4vijay/vizix

go mod init github.com/vs4vijay/vizix

go build

~/go/bin/cobra add list

createCmd.MarkFlagRequired("secret")


Logging:

log.SetFormatter(&log.TextFormatter{ForceColors: true})
log.SetOutput(colorable.NewColorableStdout())


Linting:

gofmt
golint
go vet


Git Hooks:

git config core.hooksPath .


Makefile:

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/golang/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox


Bash:

yell() { echo "FAILED> $*" >&2; }
die() { yell "$*"; exit 1; }
try() { "$@" || die "failed executing: $*"; }
log() { echo "--> $*"; }


Channels:

bye := make(chan os.Signal, 1)
signal.Notify(bye, os.Interrupt, syscall.SIGTERM)
<-bye


Context:

ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)


Mutex:

var mutex *sync.Mutex

users.mutex.Lock()
defer users.mutex.Unlock()
append(users, user)


```