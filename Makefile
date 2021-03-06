.PHONY: all

all: generate lint test

generate:
	$(info ****** GENERATE ******)
	go run ./generator/generator.go -f ./forms/data.txt
	go run ./generator/generator.go -f ./keywords/data.txt

lint:
	$(info ****** LINT ******)
	golangci-lint run
test:
	$(info ****** RUN TESTS ******)
	go test ./... -cover
