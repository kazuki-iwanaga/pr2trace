.PHONY: test
test:
	go test -v ./...

.PHONY: fix
fix:
	golangci-lint run --fix ./...

.PHONY: run
run:
	go run main.go