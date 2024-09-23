.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fix
fix:
	golangci-lint run --fix ./...

.PHONY: run
run:
	go run main.go