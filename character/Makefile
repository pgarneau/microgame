NAME:=$(shell basename $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

ifeq ($(OS),Windows_NT)
    EXT=.exe
else
    EXT=
endif

.PHONY: all build build-dist lint test build-docker run-docker

all: build build-dist lint test build-docker

build:
	go install
	go build -o bin/${NAME}${EXT}
	$(MAKE) -C migration/ build

build-dist:
	GOOS=linux GOARCH=amd64 go install
	GOOS=linux GOARCH=amd64 go build -o dist/bin/${NAME}
	$(MAKE) -C migration/ build-dist

clean: clean-dist
	-rm -rf bin/
	docker-compose down

clean-dist:
	-rm -rf dist/

lint:
#	golint -set_exit_status ./...
#	go vet

test:
#	go test -v ./...

build-docker:
	docker build -t character_go .
	$(MAKE) -C migration/ build-docker	

run-docker:
	docker-compose up
