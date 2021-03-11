.PHONY: clean

mocks:
	mockery --all

clean:
	rm -rf mocks
