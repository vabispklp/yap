test:
	$(info Testing ...)
	@go test ./...

deps:
	$(info Update deps ...)
	@go mod tidy

