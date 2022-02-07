.PHONY: test
test:
	$(info Testing ...)
	@go test ./...

.PHONY: deps
deps:
	$(info Update deps ...)
	@go mod tidy

.PHONY: postgres-up
postgres-up:
	$(info Starting postgres db ...)
	@docker-compose up -d postgres

.PHONY: postgres-stop
postgres-stop:
	$(info Stoping postgres db ...)
	@docker-compose stop postgres