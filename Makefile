.PHONY: run
run:
	go run cmd/web/main.go

.PHONY: lint
lint:
	golangci-lint run ./... -v --fix
