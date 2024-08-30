.PHONY: lint test benchmark

lint:
	@echo "Linting Go code..."
	@golangci-lint run
	@echo "Linting complete!"

test:
	@echo "Running Go tests..."
	@go test $(shell go list ./... | grep -v /faucet | grep -v /examples)
	@echo "Tests complete!"

benchmark:
	@echo "Running Go benchmarks..."
	@go test -bench=. ./...
	@echo "Benchmarks complete!"
