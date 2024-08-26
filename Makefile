.PHONY: lint

lint:
	@echo "Linting Go code..."
	@golangci-lint run
	@echo "Linting complete!"

# Set the default goal to lint (running 'make' will run 'make lint')
.DEFAULT_GOAL := lint
