dependencies:
	@go get -v -t ./...

vet:
	@go vet ./...

test:
	@go test -race .

build: test
	@go build ./...

ci: dependencies test

.PHONY: dependencies vet test
