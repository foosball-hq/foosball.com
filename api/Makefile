PROJECT ?= foosball-api
BIN ?= $(CURDIR)/bin
GOOS ?= linux
GOARCH ?= amd64

export GO111MODULE=on

.PHONY: all
all: fmt lint test build

.PHONY: build
build: 
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BIN)/$(PROJECT) $(CURDIR)/cmd/user/create.go
	chmod 755 $(BIN)/$(PROJECT)

.PHONY: test
test:
	go test $(CURDIR)/cmd/... $(CURDIR)/pkg/...

.PHONY: lint
lint:
	golint $(CURDIR)/...

.PHONY: fmt
fmt:
	gofmt -s -w $(CURDIR)

.PHONY: clean
clean:
	rm -rf $(BIN)