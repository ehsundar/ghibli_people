.PHONY: clean bench

BENCH_FLAGS ?= -benchmem -memprofile=mem.pprof -cpuprofile=cpu.pprof
MODULE_DIRS = ./internal/storage
LINTER = golangci-lint
LINTER_VERSION = v1.31.0

dependencies:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin $(LINTER_VERSION)
	go get github.com/vektra/mockery/v2/.../
	go mod download

ghp:
	go build -o ghp .

mocks:
	mockery --all

clean:
	rm -rf mocks
	rm -rf cover.out
	rm -f ghp

test: mocks
	go test ./...

cover.out: mocks
	go test -covermode=count -coverprofile=cover.out ./...
	go tool cover -html=cover.out

race: mocks
	 CGO_ENABLED=1 go test -race ./...

.PHONY: bench
BENCH ?= .
bench: mocks
	@$(foreach dir,$(MODULE_DIRS), ( \
		cd $(dir) && \
		go list . | xargs -n1 go test -bench=$(BENCH) -run=^$ $(BENCH_FLAGS) \
	) &&) true
