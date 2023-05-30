# Formatted Go files
GOFMT_FILES ?= $(shell find . -name "*.go")

default: fmt

fmt:
	gofmt -w $(GOFMT_FILES)

clean:
	go clean -testcache

neuvector:
	docker-compose up -d

test: clean
	go test ./neuvector  -v

.PHONY: \
	fmt \
	test \
	clean \
	neuvector
