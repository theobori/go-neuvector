# Formatted Go files
GOFMT_FILES ?= $(shell find . -name "*.go")
NEUVECTOR_DIR = ./neuvector
PKG_DIR = \
	./client \
	$(NEUVECTOR_DIR)/scan \
	$(NEUVECTOR_DIR)/policy \
	$(NEUVECTOR_DIR)/service

default: fmt

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	docker-compose up -d
	go test $(PKG_DIR) -v

.PHONY: \
	fmt \
	test
