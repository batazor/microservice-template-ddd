# GO TASKS =====================================================================

generate: ## Code generation
	# Generate from .go code
	@go generate ./...

	@make fmt

.PHONY: fmt
fmt: ## Format source using gofmt
	# Apply go fmt
	@gofmt -l -s -w cmd pkg internal

gosec: ## Golang security checker
	@docker run --rm -it -v $(pwd):/app -w /app/ securego/gosec:latest -exclude=G104,G110 ./...

golint: ## Linter for golang
	@docker run --rm -it -v $(pwd):/app -w /app/ golangci/golangci-lint:v1.43.0-alpine golangci-lint run ./...

test: ## Run all test
	@sh ./ops/scripts/coverage.sh

bench: ## Run benchmark tests
	go test -bench ./...
