all: test vet lint

lint:
	golint -set_exit_status ./...

vet:
	go vet ./...

test:
	go run test-client/main.go

.PHONY: lint vet test
