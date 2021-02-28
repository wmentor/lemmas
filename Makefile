.PHONY: all

all: generate test

generate:
	$(info ****** GENERATE ******)
	go generate ./forms

test:
	$(info ****** RUN TESTS ******)
	go test ./... -cover
