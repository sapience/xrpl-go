.PHONY: lint test benchmark

lint:
	@echo "Linting Go code..."
	@golangci-lint run
	@echo "Linting complete!"

test-all:
	@echo "Running Go tests..."
	@go test -v $(shell go list ./... | grep -v /faucet | grep -v /examples)
	@echo "Tests complete!"

test-binary-codec:
	@echo "Running Go tests for binary codec package..."
	@go test -v ./binary-codec/...
	@echo "Tests complete!"

test-address-codec:
	@echo "Running Go tests for address codec package..."
	@go test -v ./address-codec/...
	@echo "Tests complete!"

test-keypairs:
	@echo "Running Go tests for keypairs package..."
	@go test -v ./keypairs/...
	@echo "Tests complete!"

test-xrpl:
	@echo "Running Go tests for xrpl package..."
	@go test -v ./xrpl/...
	@echo "Tests complete!"

test-ci:
	@echo "Running Go tests..."
	@go clean -testcache
	@go test -v $(shell go list ./... | grep -v /faucet | grep -v /examples)
	@echo "Tests complete!"

benchmark:
	@echo "Running Go benchmarks..."
	@go test -bench=. ./...
	@echo "Benchmarks complete!"
