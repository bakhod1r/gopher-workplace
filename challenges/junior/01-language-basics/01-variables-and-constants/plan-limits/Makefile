.PHONY: all test vet fmt fmt-check lint verify clean

all: verify

## test: run unit tests
test:
	go test ./...

## test-v: run unit tests, verbose
test-v:
	go test -v ./...

## vet: static analysis
vet:
	go vet ./...

## fmt: format code in place
fmt:
	gofmt -w .

## fmt-check: fail if any file is unformatted
fmt-check:
	@out="$$(gofmt -l .)"; \
	if [ -n "$$out" ]; then \
		echo "unformatted files:"; echo "$$out"; exit 1; \
	fi

## lint: golangci-lint if installed, else skip
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed, skipping"; \
	fi

## verify: full gate — fmt-check + vet + test
verify: fmt-check vet test
	@echo "PASS: challenge validated"

## clean: remove go test cache for this module
clean:
	go clean -testcache
