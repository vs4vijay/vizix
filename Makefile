PROJECT 	:= "vizix"
VERSION 	:= 0.0.0

GIT_COMMIT 	:= `git rev-parse HEAD`
GIT_SHA 	:= `git rev-parse --short HEAD`
GIT_TAG 	:= `git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "canary"`
BUILD_TIME  := `date`

LDFLAGS := ""
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.SemVer=$(GIT_TAG)
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.GitCommit=$(GIT_COMMIT)
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.BuildTime=$(BUILD_TIME)

OS 			:= `uname`
OS_LIST		:= darwin freebsd linux openbsd

ARCH 		:= `uname -m`
ARCH_LIST	:= 386 amd64

.PHONY: default info clean

default: info

info:
	@echo "info..."
	@echo "Version:       		${GIT_TAG}"
	@echo "Git Commit:       	${GIT_COMMIT}"
	@echo "Git SHA:       		${GIT_SHA}"
	@echo "Build Time:       	${BUILD_TIME}"

fmt:
	gofmt -l -w .

build: info
	go build -v -ldflags "$(LDFLAGS)"

release-dry-run:
	goreleaser --snapshot --skip-publish --rm-dist

release-using-gorelease:
	goreleaser --rm-dist

clean:
	@echo "cleaning..."