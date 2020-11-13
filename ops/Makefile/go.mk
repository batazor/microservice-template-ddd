# GO TASKS =====================================================================

generate: ## Code generation
	# proto generation metadata entity
	@protoc -I/usr/local/include -I. \
		--go_out=Minternal/user/domain/user.proto=.:. \
		--go-grpc_out=Minternal/user/domain/user.proto=.:. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		internal/user/domain/user.proto

	@protoc -I/usr/local/include -I. \
		--go_out=Minternal/billing/domain/billing.proto=.:. \
		--go-grpc_out=Minternal/billing/domain/billing.proto=.:. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		internal/billing/domain/billing.proto

	# Generate from .go code
	@go generate internal/di/wire.go

	@make fmt

.PHONY: fmt
fmt: ## Format source using gofmt
	# Apply go fmt
	@gofmt -l -s -w cmd pkg internal

gosec: ## Golang security checker
	@gosec -exclude=G104,G110 ./...

golint: ## Linter for golang
	@golangci-lint run ./...

test: ## Run all test
	@sh ./ops/scripts/coverage.sh

bench: ## Run benchmark tests
	go test -bench ./...
