.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test -race -short -coverpkg=./... -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | grep total | awk '{print "Total coverage: " $$3}'
