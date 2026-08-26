.PHONY: all test test-v vet fmt fmt-check lint verify clean

all: verify

test:
	go test ./...

test-v:
	go test -v ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@out="$$(gofmt -l .)"; \
	if [ -n "$$out" ]; then \
		echo "unformatted files:"; echo "$$out"; exit 1; \
	fi

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed, skipping"; \
	fi

verify: fmt-check vet test
	@echo "PASS: challenge validated"

clean:
	go clean -testcache
