.PHONY: all

all: generate lint test

generate:
	$(info ****** GENERATE ******)
	go generate ./dicts
	go generate ./forms
	go generate ./keywords

lint:
	$(info ****** LINT ******)
	golangci-lint run
test:
	$(info ****** RUN TESTS ******)
	go test ./... -cover
