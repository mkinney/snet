.DEFAULT_GOAL := fmt

.PHONY:fmt build

fmt:
	go fmt ./...

build:
	cd cmd/tc; go build
	cd cmd/ts; go build
	cd cmd/uc; go build
	cd cmd/us; go build

clean:
	@rm cmd/tc/tc cmd/ts/ts cmd/uc/uc cmd/us/us 2> /dev/null || true

lint:
	golangci-lint run cmd/tc/tc.go
	golangci-lint run cmd/ts/ts.go
	golangci-lint run cmd/uc/uc.go
	golangci-lint run cmd/us/us.go

FORCE: ;

