.DEFAULT_GOAL := help

PHONY: build unit-test integration-test clean help

## This help screen
help:
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\_0-9%:\\]+/ { \
	helpMessage = match(lastLine, /^## (.*)/); \
	if (helpMessage) { \
	helpCommand = $$1; \
	helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
	gsub("\\\\", "", helpCommand); \
	gsub(":+$$", "", helpCommand); \
	printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
	} \
        } \
        { lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"

## Build the ethcli binary in the bin/ directory
build:
	@ printf "\nBuilding ethcli binary in bin/...\n\n"
	@ go build -o bin/ethcli main.go 2>/dev/null || printf "\nBuild failed!\n\n"

## Run unit tests
unit-test:
	@ printf "\nRunning tests...\n\n"
	@ go test -v ./...

## Run integration tests
integration-test:
	@ printf "\nRunning integration tests...\n\n"
	@ ./test/integration.sh || printf "\nIntegration tests failed!\n\n"

## Clean up the build and testing artifacts
clean:
	@ printf "\nCleaning up...\n\n"
	@ docker rm -f ganache-cli >/dev/null 2>&1 || printf "No ganache-cli container found.\n"
	@ rm -rf bin/ || printf "\nCleanup failed!\n\n"