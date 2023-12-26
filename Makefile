.DEFAULT_GOAL := fmt

.PHONY:fmt build

fmt:
	go fmt ./...

build:
	go build tc.go
	go build ts.go
	go build uc.go
	go build us.go

clean:
	@rm tc ts uc us 2> /dev/null || true

lint:
	golangci-lint run tc.go
	golangci-lint run ts.go
	golangci-lint run uc.go
	golangci-lint run us.go

FORCE: ;

