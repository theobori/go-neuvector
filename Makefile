# Formatted Go files
GOFMT_FILES ?= $(shell find . -name "*.go")
CONTROLLER_DIR = ./controller
PKG_DIR = \
	$(CONTROLLER_DIR)/policy \
	$(CONTROLLER_DIR)/admission \
	$(CONTROLLER_DIR)/scan \
	$(CONTROLLER_DIR)/service

default: fmt

fmt:
	gofmt -w $(GOFMT_FILES)

clean:
	go clean -testcache

test: clean
	docker-compose up -d
	go test ./client  -v
	go test $(CONTROLLER_DIR)/federation  -v
	@echo Idle 10 seconds for logon session timeout
	@sleep 10
	go test $(PKG_DIR) -v

.PHONY: \
	fmt \
	test \
	clean
