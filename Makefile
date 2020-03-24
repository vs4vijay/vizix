PROJECT := "vizix"
VERSION := 0.0.1
GIT_COMMIT := `git rev-parse HEAD`
GIT_SHA := `git rev-parse --short HEAD`

LDFLAGS := ""
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.Release=$(VERSION)
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.GitCommit=$(GIT_COMMIT)

.PHONY: info clean

info:
	@echo "info..."
	@echo "Version:       		${VERSION}"
	@echo "Git Commit:       	${GIT_COMMIT}"
	@echo "Git SHA:       		${GIT_SHA}"

fmt:
	gofmt -l -w .

build: info
	go build -v -ldflags "$(LDFLAGS)"

clean:
	@echo "cleaning..."