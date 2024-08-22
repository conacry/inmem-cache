.PHONY: build
build:
	go build ./...

.PHONY: test/cover
test/cover:
	go test -coverpkg=./... -coverprofile=coverage.out ./test/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.PHONY: test/run
test/run:
	go test -race ./...

.PHONY: test/clean_and_run
test/clean_and_run:
	go clean -testcache
	go test -race ./...
