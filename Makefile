.PHONY: all

all: generate test

generate:
	$(info ****** GENERATE ******)
	go generate ./forms
	go generate ./keywords

test:
	$(info ****** RUN TESTS ******)
	go test ./... -cover
