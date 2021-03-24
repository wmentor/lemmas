.PHONY: all

all: generate lint test

generate:
	$(info ****** GENERATE ******)
	go run ./generator/generator.go -f ./dicts/dict_cities.txt
	go run ./generator/generator.go -f ./dicts/dict_companies.txt
	go run ./generator/generator.go -f ./dicts/dict_countries.txt
	go run ./generator/generator.go -f ./dicts/dict_m_lastnames.txt
	go run ./generator/generator.go -f ./dicts/dict_m_names.txt
	go run ./generator/generator.go -f ./dicts/dict_w_lastnames.txt
	go run ./generator/generator.go -f ./dicts/dict_w_names.txt
	go run ./generator/generator.go -f ./forms/data.txt
	go run ./generator/generator.go -f ./keywords/data.txt -shift 1

lint:
	$(info ****** LINT ******)
	golangci-lint run
test:
	$(info ****** RUN TESTS ******)
	go test ./... -cover
