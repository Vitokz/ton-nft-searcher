.PHONY: start

SEARCHER_BINARY_NAME := searcher

test:
	go clean -testcache
	go test ./...

lint:
	golangci-lint run --fix

run-searcher: build-searcher
	./$(SEARCHER_BINARY_NAME)

build-searcher:
	go build -o $(SEARCHER_BINARY_NAME) ./cmd/

swagger-regen:
	swagger generate spec --scan-models -o ./api/swagger-ui/swagger.json -w ./cmd/
