# vizix

## Getting Started

```console
go get github.com/vs4vijay/vizix
```

---

## Development

```console

go run main.go

```

---

### Commands
```
~/go/bin/cobra init --pkg-name github.com/vs4vijay/vizix

go mod init github.com/vs4vijay/vizix

go build

~/go/bin/cobra add list

createCmd.MarkFlagRequired("secret")
```

### Logging
```
log.SetFormatter(&log.TextFormatter{ForceColors: true})
log.SetOutput(colorable.NewColorableStdout())
```
- Log Verbosity: cmd.PersistentFlags().CountVarP(&verbosity, "verbosity", "v", "set verbosity")


### Linting

- gofmt
- goimports
- golint
- go vet


### Git Hooks
```
git config core.hooksPath .
```

### Makefile
```
tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/golang/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox
```

### Channels
```
bye := make(chan os.Signal, 1)
signal.Notify(bye, os.Interrupt, syscall.SIGTERM)
<-bye
```

### Context
```
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
```

### Mutex
```
var mutex *sync.Mutex

users.mutex.Lock()
defer users.mutex.Unlock()
append(users, user)
```

### Testing
includes following suite of tests.
- `make test-lint`: runs linter/style checks
- `make test-unit`: runs basic unit tests
- `make test`: runs all of the above

### JSON
```
json.NewEncoder(writer).Encode(todos)
```

### System Exec
```
out, err := exec.Command("ls").Output()
```

### OS Detection:

- `runtime.GOOS == "windows"`


### APIs

- `net/http` - Native Implementation
- `chi` - lightweight, compatible net.Http
- `mux` - 
- `gin` - 


#### Using net/http:

```
http.Handle("/", server) // or http.HandleFunc("/", someFunc)
http.ListenAndServe(address, nil)
```
- server should have Handler interface, which should have ServeHTTP method

#### Using Mux: 

```
router := mux.NewRouter()
router.HandleFunc("/", Index)
http.ListenAndServe(address, router)
```
- To get params - `mux.Vars(request)`

### Bash:
```yell() { echo "FAILED> $*" >&2; }
die() { yell "$*"; exit 1; }
try() { "$@" || die "failed executing: $*"; }
log() { echo "--> $*"; }
```

---

### Development Notes

```

...

```