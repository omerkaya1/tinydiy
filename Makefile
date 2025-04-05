.DEFAULT_GOAL := help

.PHONY: build-blinky
build-blinky: ## build blinky
	tinygo flash --target=arduino ./cmd/blinky/main.go

.PHONY: lbtn
lbtn: ## build light button switch
	tinygo flash --target=arduino ./cmd/lbtn/main.go

.PHONY: help
help: ## print this help and exit
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
