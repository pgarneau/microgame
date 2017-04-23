NAME:=$(shell basename $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

ifeq ($(OS),Windows_NT)
    EXT=.exe
else
    EXT=
endif

.PHONY: all build build-dist

all: build build-dist

build:
	go install
	go build -o bin/${NAME}${EXT}

build-dist:
	GOOS=linux GOARCH=amd64 go install
	GOOS=linux GOARCH=amd64 go build -o dist/bin/${NAME}

build-docker:
	docker build -t pgarneau/character_migration .

clean: clean-dist
	-rm bin/${NAME}${EXT}

clean-dist:
	-rm dist/bin/${NAME}

publish-docker:
	docker push pgarneau/character_migration
