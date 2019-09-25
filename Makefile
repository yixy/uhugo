.PHONY: build clean

UHUGO_VERSION=0.0.1
UHUGO_BIN=uhugo
LDFLAGS_PATH=github.com/yixy/uhugo/cmd

GO_ENV=CGO_ENABLED=1
GO_FLAGS=-ldflags="-X $(LDFLAGS_PATH).ver=$(UHUGO_VERSION) -X '$(LDFLAGS_PATH).env=`uname -mv`' -X '$(LDFLAGS_PATH).buildTime=`date`'"
GO=env $(GO_ENV) go

UNAME := $(shell uname)

# build uhugo cli
build:
	mkdir target
	$(GO) build $(GO_FLAGS) -o target/$(UHUGO_BIN) .
# test
test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
# clean all build result
clean:
	go clean ./...
	rm -rf target
