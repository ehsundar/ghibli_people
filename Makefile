.PHONY: clean bench

BENCH_FLAGS ?= -benchmem -memprofile=mem.pprof -cpuprofile=cpu.pprof
MODULE_DIRS = ./internal/storage

mocks:
	mockery --all

clean:
	rm -rf mocks
	rm -rf cover.out
	rm

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
