# vizix

[![Release](https://github.com/vs4vijay/vizix/workflows/Release/badge.svg)](https://github.com/vs4vijay/vizix/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/vs4vijay/vizix)](https://goreportcard.com/report/github.com/vs4vijay/vizix)
[![Maintainability](https://api.codeclimate.com/v1/badges/ef8ddc6dc9dc972a7968/maintainability)](https://codeclimate.com/github/vs4vijay/vizix/maintainability)

## Getting Started

```console
go get github.com/vs4vijay/vizix
```

OR (Via **Homebrew**)

```console
brew install vs4vijay/vizix/vizix
```

OF (Via **Installer Script**)

```console
curl -fsSL https://raw.githubusercontent.com/vs4vijay/vizix/vizix/scripts/install.sh | bash
```

---

## Development

```console
go run main.go
```

- Auto Reload:
  - `Air` - https://github.com/cosmtrek/air
  - WatchMan: https://github.com/facebook/watchman
  - modd: https://github.com/cortesi/modd
  - entr: http://eradman.com/entrproject
  - https://github.com/codegangsta/gin
---

### Guidelines

- Prefer `goimports` then `gofmt`
- Use `golint` for style mistakes
- https://github.com/golang/go/wiki/CodeReviewComments

---

### Commands
```bash
~/go/bin/cobra init --pkg-name github.com/vs4vijay/vizix

go mod init github.com/vs4vijay/vizix

go build

~/go/bin/cobra add list

createCmd.MarkFlagRequired("secret")
```

---

### Logging
```golang
log.SetFormatter(&log.TextFormatter{ForceColors: true})
log.SetOutput(colorable.NewColorableStdout())
```
- Log Verbosity: `cmd.PersistentFlags().CountVarP(&verbosity, "verbosity", "v", "set verbosity")`

---

### Linting
- gofmt
- goimports
- golint
- go vet

---

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

- `make test-lint`: runs linter/style checks
- `make test-unit`: runs basic unit tests
- `make test`: runs all of the above

```

---

### I/O

```golang

s := bufio.Scanner(os.Stdin)
s.Scan()

// OR

scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
  break
}
if err := scanner.Err(); err != nil {
  panic(err)
}

// OR

serverData, err := bufio.NewReader(connection).ReadString('\n')

```

---

### JSON / Marshalling / Un-marshalling

```golang

// Marshalling 1
j, err := json.Marshal(p)

// Marshalling 2
json.NewEncoder(writer).Encode(todos)

// Un-marshalling 1
body := make(map[string]interface{})
bodyBytes, _ := ioutil.ReadAll(request.Body)
err := json.Unmarshal(bodyBytes, &body)

// Un-marshalling 2
body := make(map[string]interface{})
err := json.NewDecoder(request.Body).Decode(&body)

```

---

```golang

// Make HTTP Call

payloadBytes, err := json.Marshal(payload)
if err != nil {
    return nil, err
}

client = &http.Client{
	Timeout: 10 * time.Second,
}
request, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))

request.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
request.Header.Add("Authorization", "Bearer " + apiKey)
response, err := client.Do(request)
if err != nil {
    return nil, err
}

defer response.Body.Close()
body, err := ioutil.ReadAll(response.Body)
if err != nil {
    return nil, err
}

string(body)

// Convert Payload - I
values := map[string]string{"username": username, "password": password}
dataBytes, _ := json.Marshal(values)
resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(dataBytes))

// Convert Payload - II, Mainly for Streaming
buf := new(bytes.Buffer)
json.NewEncoder(buf).Encode(body)
resp, err := http.Post(authAuthenticatorUrl, "application/json", buf)


// Ready Body
body, err := ioutil.ReadAll(resp.Body)

// Read Body to Object
var result map[string]interface{}
json.NewDecoder(resp.Body).Decode(&result)

```

---

### Channels
```golang
bye := make(chan os.Signal, 1)
signal.Notify(bye, os.Interrupt, syscall.SIGTERM)
<-bye
```

### WaitGroups
```golang
var wg sync.WaitGroup

wg.Add(1)
wg.Done()

wg.Wait()


```

---

### Signals
- signal.Notify(s1, syscall.SIGWINCH)
- signal.Ignore(syscall.SIGINT)
- signal.Stop(s1)
```golang
sigs := make(chan os.Signal, 1)
done := make(chan bool, 1)

signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

go func() {
    sig := <-sigs
    fmt.Println(sig)
    done <- true
}()

fmt.Println("awaiting signal")
<-done
fmt.Println("exiting")
```

---

### Timers and Tickers

```golang

// Timer
timer := time.NewTimer(2 * time.Second)
<-timer.C
fmt.Println("Timer 1 fired")


// Ticker
ticker := time.NewTicker(2 * time.Second)
done := make(chan bool)

go func() {
    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            fmt.Println("Tick at", t)
        }
    }
}()

time.Sleep(10 * time.Second)
ticker.Stop()
done <- true
fmt.Println("Ticker stopped")

```

---

### Context
```golang
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
OR
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

// Example 2

ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)

go func() {
    s := bufio.Scanner(os.Stdin)
    s.Scan()
    cancel()
}()

// Example 3
time.AfterFunc(time.Second, cancel)

// OR
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel() // "deadline exceeded", but "cancel()" will cancel

// viz
func viz(ctx context.Context, d time.Duration, msg string) {
    select {
    case <- time.After(d):
        fmt.Println(msg)
    case <- ctx.Done()
        fmt.Errorf(ctx.Err())
    }
}

```

---

### Mutex
```golang
var mutex *sync.Mutex

users.mutex.Lock()
defer users.mutex.Unlock()
append(users, user)
```

---

### Testing
- Table Testing

---

### System Exec
```golang

binary, lookErr := exec.LookPath("ls")

// Spawing
out, err := exec.Command("ls").Output()


// Using Syscall
args := []string{"ls", "-a", "-l", "-h"}
env := os.Environ()
execErr := syscall.Exec(binary, args, env)
```

---

### OS Detection:

- `runtime.GOOS == "windows"`


### APIs

- `net/http` - Native Implementation
- `chi` - lightweight, compatible net.Http
- `mux` -
- `gin` -
- `iris` -
- `echo` -
- `beego`

#### Using net/http:

```golang

// Option-I:

func someFunc(w http.ResponseWriter, r *http.Request) {
    w.write([]byte("Hey"))
    // fmt.Fprintln(w, "Hey")
}

http.HandleFunc("/", someFunc)
http.ListenAndServe(address, nil)

// Option-II:

http.Handle("/", server)
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

---

### Graceful Shutdown

```golang

// Trap sigterm or interrupt to gracefully shutdown the server
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt)
signal.Notify(sigChan, os.Kill)

// Block until a signal is received.
sig := <-sigChan
logger.Printf("Got signal: %v, shutting down the server\n", sig)

// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
srv.Shutdown(ctx)

// OR

quit := make(chan os.Signal)
signal.Notify(quit, os.Interrupt)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

```

---

### Bash:
```bash
yell() { echo "FAILED> $*" >&2; }
die() { yell "$*"; exit 1; }
try() { "$@" || die "failed executing: $*"; }
log() { echo "--> $*"; }
```

---

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

---

### Build Binary

- Build Tags: `// +build pro`
- go build -tags pro

---

### Distribute Binaries

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

---

### Deployment
- Fly:
- OpenShift: https://manage.openshift.com/

---

### Badges
- ![Release](https://github.com/vs4vijay/vizix/workflows/Release/badge.svg) - `![Release](https://github.com/vs4vijay/vizix/workflows/Release/badge.svg)`

---

### 3rd Party Integrations
- Renovate Bot
- HoundCI
- Kodiak
- https://codecov.io/
- https://codeclimate.com/
- https://github.com/NickeManarin/ScreenToGif
- Sentry

---

### Git Hooks
```bash
git config core.hooksPath .
```
- commit - go mod tidy

---

### Development Notes

```

GO111MODULE=on
GOPROXY=https://gocenter.io
CGO_ENABLED=0

GOARCH=wasm GOOS=js go build -o app.wasm

wget: wget -q -O - https://raw.githubusercontent.com/rancher/k3d/master/install.sh | TAG=v1.3.4 bash
curl: curl -s https://raw.githubusercontent.com/rancher/k3d/master/install.sh | TAG=v1.3.4 bash

curl -sfL https://get.k3s.io | sh -

```
