.DEFAULT_GOAL := fmt

.PHONY:fmt build

fmt:
	go fmt ./...

build:
	cd tc; go build
	cd ts; go build
	cd uc; go build
	cd us; go build

clean:
	@rm tc/tc ts/ts uc/uc us/us 2> /dev/null || true

lint:
	golangci-lint run tc/tc.go
	golangci-lint run ts/ts.go
	golangci-lint run uc/uc.go
	golangci-lint run us/us.go

FORCE: ;

