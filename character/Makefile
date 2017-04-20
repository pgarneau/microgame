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

build-dist:
	GOOS=linux GOARCH=amd64 go install
	GOOS=linux GOARCH=amd64 go build -o dist/bin/${NAME}

clean: clean-dist
	-rm bin/${NAME}${EXT}

clean-dist:
	-rm dist/bin/${NAME}

lint:
#	golint -set_exit_status ./...
#	go vet

test:
#	go test -v ./...

build-docker:
	docker build -t ${NAME} .

run-docker:
	docker run --publish 6060:8080 --name ${NAME} --rm ${NAME}
