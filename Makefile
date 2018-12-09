
SHELL := /bin/bash
CURRENT_PATH = $(shell pwd)
APP_NAME = packr
APP_VERSION = "0.1.0"

# build with verison infos
VERSION_DIR = "github.com/overbool/${APP_NAME}"
BUILD_DATE = $(shell date +%FT%T)
GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
LDFLAGS = "-w -X ${VERSION_DIR}.BuildDate=${BUILD_DATE} -X ${VERSION_DIR}.CurrentCommit=${GIT_COMMIT}"

install: clean
	go install -ldflags=${LDFLAGS} ./cmd/${APP_NAME}

release: clean
	sh scripts/release.sh ${APP_NAME} ${APP_VERSION} ${LDFLAGS} ${CURRENT_PATH}/cmd/${APP_NAME}

test:
	go test -v .

clean:
	go clean
	rm -f bin/*

testcase: install
	cd cmd/packer && packer ../../test
