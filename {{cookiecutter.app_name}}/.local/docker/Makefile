APP := $(shell basename $(CURDIR))
VERSION := $(shell git describe --tags --always --dirty)
REPORTS_DIR=.reports

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
#.PHONY: bin/$(APP) bin clean start test

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# OTEL TASKS
setup: ## Setup otel docker
	docker-compose -f .local/docker/docker-compose.yaml up -d

otel-collector-logs: ## View otel collector logs
	docker-compose -f .local/docker/docker-compose.yaml logs otel-collector  -f

adot-collector-logs: ## View adot collector logs
	docker-compose -f .local/docker/docker-compose.yaml logs adot-collector  -f

status: ## View otel docker status
	docker-compose -f .local/docker/docker-compose.yaml ps

teardown: ## Teardown otel docker
	docker-compose -f .local/docker/docker-compose.yaml down



