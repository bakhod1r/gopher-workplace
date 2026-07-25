# Gopher Workplace — root Makefile
# Runs targets across every challenge module (each has its own go.mod).

CHALLENGES := $(shell find challenges -name go.mod -exec dirname {} \;)

.PHONY: help verify test list

help:
	@echo "Gopher Workplace"
	@echo ""
	@echo "  make list      list all challenges"
	@echo "  make verify    run 'make verify' in every challenge"
	@echo "  make test      run tests in every challenge"
	@echo ""
	@echo "Single challenge:  make -C challenges/junior/01-slice-dedupe verify"

## list: show all challenge modules
list:
	@for d in $(CHALLENGES); do echo "  $$d"; done

## verify: gate every challenge
verify:
	@for d in $(CHALLENGES); do \
		echo "==> $$d"; \
		$(MAKE) -C $$d verify || exit 1; \
	done

## test: test every challenge
test:
	@for d in $(CHALLENGES); do \
		echo "==> $$d"; \
		$(MAKE) -C $$d test || exit 1; \
	done
