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
```bash
~/go/bin/cobra init --pkg-name github.com/vs4vijay/vizix

go mod init github.com/vs4vijay/vizix

go build

~/go/bin/cobra add list

createCmd.MarkFlagRequired("secret")
```

### Logging
```golang
log.SetFormatter(&log.TextFormatter{ForceColors: true})
log.SetOutput(colorable.NewColorableStdout())
```
- Log Verbosity: `cmd.PersistentFlags().CountVarP(&verbosity, "verbosity", "v", "set verbosity")`


### Linting
- gofmt
- goimports
- golint
- go vet


### Git Hooks
```bash
git config core.hooksPath .
```


### Makefile
```bash
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
```golang
bye := make(chan os.Signal, 1)
signal.Notify(bye, os.Interrupt, syscall.SIGTERM)
<-bye
```


### Context
```golang
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
OR
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
```


### Mutex
```golang
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
```golang
json.NewEncoder(writer).Encode(todos)
```


### System Exec
```golang
out, err := exec.Command("ls").Output()
```


### OS Detection:

- `runtime.GOOS == "windows"`


### APIs

- `net/http` - Native Implementation
- `chi` - lightweight, compatible net.Http
- `mux` -
- `gin` -
- `iris` -
- `echo` -

#### Using net/http:

```golang
http.Handle("/", server) // or http.HandleFunc("/", someFunc)
http.ListenAndServe(address, nil)
```
- server should have Handler interface, which should have ServeHTTP method

#### Using Mux:

```golang
router := mux.NewRouter()
router.HandleFunc("/", Index)
http.ListenAndServe(address, router)
```
- To get params - `mux.Vars(request)`


### Bash:
```bash
yell() { echo "FAILED> $*" >&2; }
die() { yell "$*"; exit 1; }
try() { "$@" || die "failed executing: $*"; }
log() { echo "--> $*"; }
```


### Dockerization:

```bash
docker build -t vs4vijay/vizix .
docker run vs4vijay/vizix
```

- Cleanup:
```bash
docker container prune
docker image prune
docker network prune
docker volume prune
```

- List Dangling Images: `docker images -f dangling=true`
  - `docker rmi $(docker images --filter "dangling=true" -q --no-trunc)`
- Remove Volumes of Dangling Images: `docker volume rm $(docker volume ls -qf dangling=true)`
- Remove Containers: `docker rm $(docker ps -qa --no-trunc --filter "status=exited")`
- Remove Everything: `docker system prune -a --volumes`
- Kill All Running Containers: `docker kill $(docker ps -q)`


### Build and Distribute

- Manual Build
```bash
GOOS=darwin GOARCH=amd64 go build
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=386 go build
```

- Git tags
  - Create Tag: `git tag -a v0.0.0 -m "Initial release"`
  - Push Tag: `git push origin v0.0.0`
  - Delete Tag: `git push origin :v0.0.1`

- Go Releaser
  - `brew install goreleaser/tap/goreleaser`
  - `goreleaser init`
  - Test: `goreleaser --snapshot --skip-publish --rm-dist`
  - Release: `goreleaser --rm-dist`
  - CI/CD with Github Actions:
```yaml
name: Release

on:
  push:
    tags:
      - "v*.*.*"

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      -
        name: Checkout
        uses: actions/checkout@v2
      -
        name: Unshallow
        run: git fetch --prune --unshallow
      -
        name: Set up Go
        uses: actions/setup-go@v1
        with:
          go-version: 1.14.x
      -
        name: GoReleaser Action
        uses: goreleaser/goreleaser-action@v1.3.1
        with:
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GORELEASER_TOKEN }}
```

- Brew formula
  - Repo: [vs4vijay/homebrew-vizix](https://github.com/vs4vijay/homebrew-vizix)
```
brew install vizix
brew info vizix
brew reinstall vizix --force
brew update
brew upgrade vizix
```

### Deployment

- Fly:
- OpenShift: https://manage.openshift.com/


### Badges

- ![Release](https://github.com/srijanone/vega/workflows/Release/badge.svg) - `![Release](https://github.com/srijanone/vega/workflows/Release/badge.svg)`
-

---

### Development Notes

```

GO111MODULE=on
GOPROXY=https://gocenter.io
CGO_ENABLED=0

```
