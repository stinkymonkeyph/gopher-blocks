
.PHONY: test run

test:
	go test -v ./...

run-main:
	go run .
