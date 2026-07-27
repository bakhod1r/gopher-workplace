# Gopher Workplace — root Makefile
# Runs targets across every challenge module (each has its own go.mod),
# plus setup/run targets for the web UI + local runner.

CHALLENGES  := $(shell find challenges -name go.mod -exec dirname {} \;)
RUNNER_DIR  := site/cmd/localrunner
RUNNER_PORT ?= 7070
WEB_PORT    ?= 8080

.PHONY: help list verify test site-test \
        setup check-tools deps catalog update \
        runner serve dev site-build reconstruct clean

help:
	@echo "Gopher Workplace"
	@echo ""
	@echo "Setup"
	@echo "  make setup                 check tools + download deps + build catalog"
	@echo "  make update                git pull --ff-only, then re-run setup"
	@echo ""
	@echo "Run"
	@echo "  make dev                   runner + web UI on http://localhost:$(RUNNER_PORT)"
	@echo "                             [RUNNER_PORT=$(RUNNER_PORT)]"
	@echo "  make serve                 static UI only, no backend [WEB_PORT=$(WEB_PORT)]"
	@echo "  make catalog               regenerate site problem catalog from challenges/"
	@echo ""
	@echo "Challenges"
	@echo "  make list                  list all challenge modules"
	@echo "  make verify                fmt-check + vet + test in every challenge"
	@echo "  make test                  run tests in every challenge"
	@echo "  make site-test             test the runner + catalog generator"
	@echo "  make reconstruct           regenerate the challenge grid (idempotent)"
	@echo "  make clean                 drop test caches + built site server"
	@echo ""
	@echo "Single challenge:  make -C challenges/<level>/<topic>/<subtopic>/<name> verify"

## list: show all challenge modules
list:
	@for d in $(CHALLENGES); do echo "  $$d"; done

## verify: gate every challenge
verify:
	@for d in $(CHALLENGES); do \
		echo "==> $$d"; \
		$(MAKE) -C $$d verify || exit 1; \
	done

## site-test: gate the tooling (local runner + catalog generator)
site-test:
	@for d in $(RUNNER_DIR) site/cmd/gencatalog; do \
		echo "==> $$d"; \
		( cd $$d && gofmt -l . | grep . && exit 1 || true ) && \
		( cd $$d && go vet ./... && go test ./... ) || exit 1; \
	done

## test: test every challenge
test:
	@for d in $(CHALLENGES); do \
		echo "==> $$d"; \
		$(MAKE) -C $$d test || exit 1; \
	done

# ---------------------------------------------------------------- setup ----

## check-tools: fail early if the Go toolchain is missing or too old
check-tools:
	@command -v go >/dev/null || { echo "missing: go (https://go.dev/dl/)"; exit 1; }
	@have=$$(go env GOVERSION | sed 's/^go//'); \
	need=$$(sed -n 's/^go //p' $(RUNNER_DIR)/go.mod | head -1); \
	[ "$$(printf '%s\n%s\n' "$$need" "$$have" | sort -V | head -1)" = "$$need" ] || \
		{ echo "go $$have too old — need $$need+ (https://go.dev/dl/)"; exit 1; }
	@echo "==> go      $$(go version)"

## deps: download the local runner's module deps
deps:
	@echo "==> downloading deps ($(RUNNER_DIR))"
	@cd $(RUNNER_DIR) && go mod download

## catalog: regenerate site/web/assets/js/problems.js from challenges/
catalog:
	@bash site/scripts/build.sh

site-build: catalog

## setup: one-shot project bootstrap
setup: check-tools deps catalog
	@echo ""
	@echo "==> setup done. next:  make dev"

## update: pull the new version and re-bootstrap
update:
	@git diff --quiet || { echo "working tree dirty — commit or stash first"; exit 1; }
	@echo "==> pulling"
	@git pull --ff-only
	@$(MAKE) --no-print-directory setup
	@echo "==> updated. your solve history in ~/.gopher-workplace/runner.db is untouched"

# ------------------------------------------------------------------ run ----

## runner: start the local Go-toolchain runner (serves :$(RUNNER_PORT))
runner:
	@echo "==> runner + web UI on http://localhost:$(RUNNER_PORT)"
	@cd $(RUNNER_DIR) && GW_ROOT=$(CURDIR) go run . -port $(RUNNER_PORT)

## serve: static UI only, no backend (frontend work; needs python3)
serve:
	@bash site/scripts/serve.sh $(WEB_PORT)

## dev: runner + web UI in one process, one origin
dev: runner

## reconstruct: regenerate the challenge grid (keeps authored puzzles)
reconstruct:
	@bash scripts/reconstruct.sh

## clean: drop test caches + built site server binary
clean:
	@for d in $(CHALLENGES); do $(MAKE) --no-print-directory -C $$d clean 2>/dev/null || true; done
	@go clean -testcache 2>/dev/null || true
	@echo "==> clean"
