PROJECT 	:= "vizix"
VERSION 	:= 0.0.1

GIT_COMMIT 	:= `git rev-parse HEAD`
GIT_SHA 	:= `git rev-parse --short HEAD`
GIT_TAG 	:= `git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "canary"`
GIT_STATE	:= $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

LDFLAGS := ""
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.Release=$(VERSION)
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.GitCommit=$(GIT_COMMIT)
LDFLAGS += -X=github.com/vs4vijay/vizix/pkg/version.GitTreeState=$(GIT_STATE)

OS 			:= `uname`
OS_LIST		:= darwin freebsd linux openbsd

ARCH 		:= `uname -m`
ARCH_LIST	:= 386 amd64

.PHONY: default info clean

default: info

info:
	@echo "info..."
	@echo "Version:       		${VERSION}"
	@echo "Git Commit:       	${GIT_COMMIT}"
	@echo "Git SHA:       		${GIT_SHA}"
	@echo "Git Tag:       		${GIT_TAG}"
	@echo "Git State:       	${GIT_STATE}"

fmt:
	gofmt -l -w .

build: info
	go build -v -ldflags "$(LDFLAGS)"

release-dry-run:
	goreleaser --snapshot --skip-publish --rm-dist

clean:
	@echo "cleaning..."