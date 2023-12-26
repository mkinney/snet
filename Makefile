.DEFAULT_GOAL := fmt

.PHONY:fmt build

fmt:
	go fmt ./...

build:
	go build

clean:
	@rm snet 2> /dev/null || true

lint:
	golangci-lint run

FORCE: ;

