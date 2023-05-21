# Formatted Go files
GOFMT_FILES ?= $(shell find . -name "*.go")
PKG_DIR = ./neuvector

default: fmt

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	docker-compose up -d
	go test $(PKG_DIR) -v

.PHONY: \
	fmt \
	test
