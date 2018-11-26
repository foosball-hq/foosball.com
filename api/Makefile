PROJECT ?= foosball-api
BIN ?= $(CURDIR)/bin
GOOS ?= linux
GOARCH ?= amd64

.PHONY: all
all: fmt lint test build

.PHONY: build
build:
	./scripts/build.sh
	
.PHONY: test
test:
	go test $(CURDIR)/cmd/... $(CURDIR)/foosball/...

.PHONY: lint
lint:
	golint $(CURDIR)/...

.PHONY: fmt
fmt:
	gofmt -s -w $(CURDIR)

.PHONY: clean
clean:
	rm -rf $(BIN)