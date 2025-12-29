.PHONY: test

test:
	mkdir -p unit-test-output
	go fmt ./...
	go vet -p=$(shell nproc) ./...
	go tool gotest.tools/gotestsum --junitfile unit-test-output/report.xml --format testname -- -race -p $(shell nproc) -parallel $(shell echo $$(($(shell nproc) / 2))) -coverpkg=./... -coverprofile=unit-test-output/coverage.txt -covermode=atomic ./...
	go tool cover -func unit-test-output/coverage.txt
	go tool cover -html=unit-test-output/coverage.txt -o unit-test-output/coverage.html
	go tool github.com/boumenot/gocover-cobertura < unit-test-output/coverage.txt > unit-test-output/coverage.xml