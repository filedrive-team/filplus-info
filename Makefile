.PHONY: build

VERSION=./buildversion.go
Major=1
Minjor=1
Patch=0

TARGET=./filplus-info
SUB_DIR = frontend

build: backend $(SUB_DIR)

frontend:
	@+make -C $@

backend: clean buildversion.go
	go fmt
	go build -o ${TARGET} .


run: build
	${TARGET}

buildversion.go:
	echo package main > ${VERSION}
	echo "const Major = \"${Major}\"" >> ${VERSION}
	echo "const Minjor = \"${Minjor}\"" >> ${VERSION}
	echo "const Patch = \"${Patch}\"" >> ${VERSION}
	echo "const BuildVersion = \"`git --no-pager log --pretty="%h" -n 1`\"" >> ${VERSION}

.PHONY: clean  $(SUB_DIR)
clean:
	-rm -f ${TARGET}
	-rm -f buildversion.go
